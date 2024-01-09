package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleModel struct {
	Model
	WhoEdit
	RoleUuid string `gorm:"not null;unique;comment:角色uuid" json:"roleUuid"`
	ParentId int64  `gorm:"comment:父角色id;default:0" json:"parentId"`
	Name     string `gorm:"not null;comment:角色名" json:"name"`
	Intro    string `gorm:"comment:角色介绍" json:"intro"`
}

func (u *RoleModel) TableName() string {
	return "role"
}

func (u *RoleModel) BeforeCreate(tx *gorm.DB) (err error) {
	u.RoleUuid = uuid.NewString()
	u.CreatedAtInt64 = tx.NowFunc().Unix()
	u.UpdatedAtInt64 = tx.NowFunc().Unix()
	return
}

func (u *RoleModel) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAtInt64 = tx.NowFunc().Unix()
	return
}
