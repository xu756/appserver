package db

import (
	"context"
	"server/ent"
	"server/ent/user"
	"server/internal/xerr"
)

type dbUserModel interface {
	FindUserByUsername(ctx context.Context, username string) (userInfo *ent.User, err error)
	FindUserByMobile(ctx context.Context, mobile string) (userInfo *ent.User, err error)
	CreateUser(ctx context.Context, username, password, mobile string, creator int64) (userInfo *ent.User, err error)
}

// FindUserByUsername 根据用户名查找用户
func (m *customModel) FindUserByUsername(ctx context.Context, username string) (userInfo *ent.User, err error) {
	tx, err := m.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	userInfo, err = tx.User.Query().
		Where(user.Username(username), user.Deleted(false)).First(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, xerr.ErrMsg(xerr.UserNotExist)
	case err != nil:
		return nil, xerr.DbFindErr()
	default:
		return userInfo, nil
	}
}

// FindUserByMobile 根据手机号查找用户
func (m *customModel) FindUserByMobile(ctx context.Context, mobile string) (userInfo *ent.User, err error) {
	tx, err := m.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	userInfo, err = tx.User.Query().
		Where(user.Mobile(mobile), user.Deleted(false)).First(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, xerr.ErrMsg(xerr.UserNotExist)
	case err != nil:
		return nil, xerr.DbFindErr()
	default:
		return userInfo, nil
	}
}

// CreateUser 创建用户
func (m *customModel) CreateUser(ctx context.Context, username, password, mobile string, creator int64) (userInfo *ent.User, err error) {
	tx, err := m.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	userInfo, err = tx.User.Create().
		SetUsername(username).
		SetPassword(password).
		SetMobile(mobile).
		SetCreator(creator).
		SetEditor(creator).
		SetVersion(1).
		Save(ctx)
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return userInfo, err
}
