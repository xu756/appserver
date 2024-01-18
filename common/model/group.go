package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	Model
	GroupUuid string `gorm:"not null;unique;comment:组uuid" json:"groupUuid"`
	ParentId  int64  `gorm:"comment:父组id;default:0" json:"parentId"`
	Name      string `gorm:"not null;comment:组名" json:"name"`
	Intro     string `gorm:"comment:组介绍" json:"intro"`
}

func (g *Group) TableName() string {
	return "group"
}

func (g *Group) BeforeCreate(tx *gorm.DB) (err error) {
	g.GroupUuid = uuid.NewString()
	g.CreatedAtInt64 = tx.NowFunc().Unix()
	g.UpdatedAtInt64 = tx.NowFunc().Unix()
	return
}

func (g *Group) BeforeUpdate(tx *gorm.DB) (err error) {
	g.UpdatedAtInt64 = tx.NowFunc().Unix()
	return
}
