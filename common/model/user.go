package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserInfo struct {
	UserUuid string `gorm:"not null;unique;comment:用户uuid" json:"userUuid"` // 用户uuid
	UserName string `gorm:"not null;comment:用户名" json:"userName"`           // 用户名
	Avatar   string `gorm:"comment:头像" json:"avatar"`                       // 头像
	Mobile   string `gorm:"not null;unique;comment:手机号" json:"mobile"`      // 手机号
}

type UserModel struct {
	Model
	WhoEdit
	UserInfo
	Password string `gorm:"default:'123456';comment:密码" json:"passWord"` // 密码
}

func (u *UserModel) TableName() string {
	return "user"
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	u.UserUuid = uuid.NewString()
	u.CreatedAtInt64 = tx.NowFunc().Unix()
	u.UpdatedAtInt64 = tx.NowFunc().Unix()
	return
}

func (u *UserModel) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAtInt64 = tx.NowFunc().Unix()
	return
}
