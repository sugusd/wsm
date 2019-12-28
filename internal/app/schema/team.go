package schema

import "github.com/axetroy/terminal/internal/app/db"

// 团队信息相关
type TeamPure struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	OwnerID string  `json:"owner_id"`
	Remark  *string `json:"remark"`
}

type Team struct {
	TeamPure
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TeamStat struct {
	Team
	MemberNum int `json:"member_num"` // 团队的成员数量
	HostNum   int `json:"host_num"`   // 拥有的服务器数量
}

// 团队成员信息
type TeamMember struct {
	ProfilePublic
	Role      db.TeamRole `json:"role"`       // 用户在团队的角色
	CreatedAt string      `json:"created_at"` // 用户加入团队的时间
}
