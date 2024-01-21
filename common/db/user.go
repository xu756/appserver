package db

import (
	"context"
	"log"
	"server/ent"
	"server/ent/user"
	"server/internal/xerr"
)

type dbUserModel interface {
	FindUserByUsername(ctx context.Context, username string) (userInfo *ent.User, err error)
	FindUserByMobile(ctx context.Context, mobile string) (userInfo *ent.User, err error)
	CreateUser(ctx context.Context, username, password, mobile string, creator int64) (userInfo *ent.User, err error)
	UpdateUserAvatar(ctx context.Context, uuid string, avatar string, editor int64) (userInfo *ent.User, err error)
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
		return userInfo, tx.Commit()
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
		return userInfo, tx.Commit()
	}
}

// CreateUser 创建用户
func (m *customModel) CreateUser(ctx context.Context, username, password, mobile string, creator int64) (userInfo *ent.User, err error) {
	tx, err := m.client.Tx(ctx)
	if err != nil {
		log.Print(err)
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
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return userInfo, tx.Commit()
}

// UpdateUserAvatar 更新用户头像
func (m *customModel) UpdateUserAvatar(ctx context.Context, uuid string, avatar string, editor int64) (userInfo *ent.User, err error) {
	tx, err := m.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	id, err := tx.User.Update().
		SetAvatar(avatar).
		SetEditor(editor).
		Where(user.UUID(uuid)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	userInfo, err = tx.User.Get(ctx, int64(id))
	if err != nil {
		return nil, err
	}
	return userInfo, tx.Commit()
}
