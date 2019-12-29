[![Build Status](https://github.com/axetroy/wsm/workflows/backen/badge.svg)](https://github.com/axetroy/wsm/actions)
[![Build Status](https://github.com/axetroy/wsm/workflows/frontend/badge.svg)](https://github.com/axetroy/wsm/actions)
[![Coverage Status](https://coveralls.io/repos/github/axetroy/wsm/badge.svg?branch=master)](https://coveralls.io/github/axetroy/wsm?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/wsm)](https://goreportcard.com/report/github.com/axetroy/wsm)
[![DeepScan grade](https://deepscan.io/api/teams/6484/projects/8581/branches/105883/badge/grade.svg)](https://deepscan.io/dashboard#view=project&tid=6484&pid=8581&bid=105883)
![License](https://img.shields.io/github/license/axetroy/wsm.svg)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/wsm.svg)

## Web Server Manager

通过 Web 来管理远端服务器

管理员邀请成员加入团队，管理服务器。

成员无需关心服务器的`地址/帐号/密码/密钥`等敏感信息，即可连接服务器进行操作

在成员离职后，在团队中移除成员即可

## 使用

```shell
$ git clone https://github.com/axetroy/wsm.git $GOPATH/src/github.com/axetroy/wsm
$ cd $GOPATH/src/github.com/axetroy/wsm
```

### 启动后端 API

```shell
$ go run cmd/user/main.go start
```

### 启动前端页面

```shell
$ cd ./frontend
$ yarn
$ npm run dev
```

## 部署

TODO: 在往后会打包成 Docker 镜像进行部署

## 技术栈

- Golang
- Node.js + Nuxt

## TODO

- [ ] 一次性分享终端
- [ ] 终端操作记录
- [ ] 操作记录回放

## 许可协议

[Apache License 2.0](LICENSE)
