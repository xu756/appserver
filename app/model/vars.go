package model

import (
	"github.com/xu756/appserver/internal/xerr"
)

var ErrNotFound = xerr.DatabaseErr("未找到数据")
