package model

import "gorm.io/gorm"

type UserGroup struct {
	Model
	UserUuid  string `gorm:"not null;comment:用户uuid" json:"userUuid"` // 用户uuid
	UserId    int64  `gorm:"not null;comment:用户id" json:"userId"`     // 用户id
	GroupUuid string `gorm:"not null;comment:组uuid" json:"groupUuid"` // 组uuid
	GroupId   int64  `gorm:"not null;comment:组id" json:"groupId"`     // 组id
}

func (u *UserGroup) TableName() string {
	return "user_group"
}

func (u *UserGroup) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAtInt64 = tx.NowFunc().Unix()
	u.UpdatedAtInt64 = tx.NowFunc().Unix()
	return
}

func (u *UserGroup) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAtInt64 = tx.NowFunc().Unix()
	return
}
