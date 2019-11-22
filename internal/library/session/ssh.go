package session

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

type Terminal struct {
	client  *ssh.Client
	session *ssh.Session
	config  Config
	exitMsg string
	stdout  io.Reader
	stdin   io.Writer
	stderr  io.Reader
}

type Config struct {
	Username string
	Host     string
	Port     uint
	Password string
	Width    int // pty width
	Height   int // pty height
}

func (t *Terminal) updateTerminalSize() {
	go func() {
		// SIGWINCH is sent to the process when the window size of the terminal has
		// changed.
		sigwinchCh := make(chan os.Signal, 1)
		signal.Notify(sigwinchCh, syscall.SIGWINCH)

		fd := int(os.Stdin.Fd())
		termWidth, termHeight, err := terminal.GetSize(fd)
		if err != nil {
			fmt.Println(err)
		}

		for {
			select {
			// The client updated the size of the local PTY. This change needs to occur
			// on the server side PTY as well.
			case sigwinch := <-sigwinchCh:
				if sigwinch == nil {
					return
				}
				currTermWidth, currTermHeight, err := terminal.GetSize(fd)

				if err != nil {
					fmt.Printf("Unable get terminal size: %s", err)
					continue
				}

				// Terminal size has not changed, don't do anything.
				if currTermHeight == termHeight && currTermWidth == termWidth {
					continue
				}

				if err := t.session.WindowChange(currTermHeight, currTermWidth); err != nil {
					fmt.Printf("Unable to send window-change reqest: %s.", err)
					continue
				}

				termWidth, termHeight = currTermWidth, currTermHeight
			}
		}
	}()
}

func (t *Terminal) Close() error {
	if err := t.session.Close(); err != nil {
		return err
	}

	return t.client.Close()
}

func (t *Terminal) Connect(stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	defer func() {
		if t.exitMsg == "" {
			//_, _ = fmt.Fprintln(stdout, "the connection was closed on the remote side on ", time.Now().Format(time.RFC822))
		} else {
			_, _ = fmt.Fprintln(stdout, t.exitMsg)
		}
	}()

	fd := int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fd)

	if err != nil {
		return err
	}

	defer terminal.Restore(fd, state)

	termType := os.Getenv("TERM")

	if termType == "" {
		termType = "xterm-256color"
	}

	if err = t.session.RequestPty(termType, t.config.Height, t.config.Width, ssh.TerminalModes{}); err != nil {
		return err
	}

	//t.updateTerminalSize()

	if t.stdin, err = t.session.StdinPipe(); err != nil {
		return err
	}

	if t.stdout, err = t.session.StdoutPipe(); err != nil {
		return err
	}

	if t.stderr, err = t.session.StderrPipe(); err != nil {
		return err
	}

	go io.Copy(stderr, t.stderr)
	go io.Copy(stdout, t.stdout)

	go func() {
		buf := make([]byte, 128)
		for {
			n, err := stdin.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			if n > 0 {
				_, err = t.stdin.Write(buf[:n])
				if err != nil {
					fmt.Println(err)
					t.exitMsg = err.Error()
					return
				}
			}
		}
	}()

	if err = t.session.Shell(); err != nil {
		return err
	}

	quit := make(chan int)
	go func() {
		if err = t.session.Wait(); err != nil {
			return
		}
		quit <- 1
	}()

	return nil
}

func NewTerminal(config Config) (*Terminal, error) {
	sshConfig := &ssh.ClientConfig{
		User: config.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		BannerCallback:  ssh.BannerDisplayStderr(),
	}

	client, err := ssh.Dial("tcp", net.JoinHostPort(config.Host, string(config.Port)), sshConfig)

	if err != nil {
		return nil, err
	}

	session, err := client.NewSession()

	if err != nil {
		return nil, err
	}

	s := Terminal{
		client:  client,
		config:  config,
		session: session,
	}

	return &s, nil
}
