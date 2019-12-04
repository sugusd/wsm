package db

import (
	"time"

	"github.com/axetroy/terminal/internal/library/util"
	"github.com/jinzhu/gorm"
)

type HostRecordType string

const (
	HostRecordTypeOwner        HostRecordType = "owner"
	HostRecordTypeCollaborator HostRecordType = "collaborator"
)

type HostRecord struct {
	Id     string         `gorm:"primary_key;not null;unique;index;type:varchar(32);" json:"id"` // 记录 ID
	UserID string         `gorm:"not null;index;type:varchar(32);" json:"user_id"`               // 对应的用户 ID
	HostID string         `gorm:"not null;index;type:varchar(32);" json:"host_id"`               // 对应的服务器 ID
	Host   Host           `gorm:"foreignkey:HostID" json:"host"`                                 // **关联外键**
	Type   HostRecordType `gorm:"not null;index;type:varchar(32);" json:"type"`                  // 类型

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (u *HostRecord) TableName() string {
	return "host_record"
}

func (u *HostRecord) BeforeCreate(scope *gorm.Scope) error {
	// 生成ID
	uid := util.GenerateId()
	if err := scope.SetColumn("id", uid); err != nil {
		return err
	}

	return nil
}
