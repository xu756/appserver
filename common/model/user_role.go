package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoleModel struct {
	Model
	WhoEdit
	UserUuid string `gorm:"not null;comment:用户uuid" json:"userUuid"` // 用户uuid
	UserId   int64  `gorm:"not null;comment:用户id" json:"userId"`     // 用户id
	RoleUuid string `gorm:"not null;comment:角色uuid" json:"roleUuid"` // 角色uuid
	RoleId   int64  `gorm:"not null;comment:角色id" json:"roleId"`     // 角色id
}

func (u *UserRoleModel) TableName() string {
	return "user_role"
}

func (u *UserRoleModel) BeforeCreate(tx *gorm.DB) (err error) {
	u.UserUuid = uuid.NewString()
	u.CreatedAtInt64 = tx.NowFunc().Unix()
	u.UpdatedAtInt64 = tx.NowFunc().Unix()
	return
}

func (u *UserRoleModel) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAtInt64 = tx.NowFunc().Unix()
	return
}
