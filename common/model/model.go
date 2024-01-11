package model

import (
	"database/sql"
	"time"
)

type Model struct {
	ID        int64 `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

type WhoEdit struct {
	CreatedAtInt64 int64  `gorm:"-" json:"createdAtInt64"`
	UpdatedAtInt64 int64  `gorm:"-" json:"updatedAtInt64"`
	Creator        uint64 `gorm:"not null;comment:创建者" json:"creator"`   // 创建者
	Editor         uint64 `gorm:"not null;comment:编辑者" json:"editor"`    // 编辑者
	Deleted        bool   `gorm:"not null;default:false" json:"deleted"` // 是否删除
}
