package model

import (
	"github.com/xu756/appserver/internal/xerr"
)

var ErrNotFound = xerr.DbErr("未找到本条数据")
