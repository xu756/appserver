package db

import (
	"context"
	"github.com/google/uuid"
	"server/common/model"
	"server/internal/xerr"
)

type dbUserModel interface {
	FindUserByUsername(ctx context.Context, userName string) (user *model.User, err error)
	FindUserByMobile(ctx context.Context, mobile string) (user *model.User, err error)
	CreateUser(ctx context.Context, userName, mobile string, creator uint64) (err error)
}

// FindUserByUsername 根据用户名查找用户
func (m *customModel) FindUserByUsername(ctx context.Context, userName string) (user *model.User, err error) {
	tx := m.Db.WithContext(ctx).Where("deleted = ?", false)
	err = tx.Model(&model.User{}).Where("user_name = ?", userName).Limit(1).Find(&user).Error
	switch {
	case user.UserUuid == "":
		return nil, xerr.DbErr(xerr.UserNotExist)
	case err != nil:
		return nil, xerr.DbErr(xerr.DbFindErr)
	}
	return user, nil
}

// FindUserByMobile 根据手机号查找用户
func (m *customModel) FindUserByMobile(ctx context.Context, mobile string) (user *model.User, err error) {
	tx := m.Db.WithContext(ctx).Where("deleted = ?", false)
	err = tx.Model(&model.User{}).Where("mobile = ?", mobile).Limit(1).Find(&user).Error
	switch {
	case user.UserUuid == "":
		return nil, xerr.DbErr(xerr.UserMobileNotExist)
	case err != nil:
		return nil, xerr.DbErr(xerr.DbFindErr)
	}
	return user, nil
}

// CreateUser 创建用户
func (m *customModel) CreateUser(ctx context.Context, userName, mobile string, creator uint64) (err error) {
	tx := m.Db.WithContext(ctx).Model(&model.User{})
	var user *model.User
	user.UserUuid = uuid.NewString()
	user.UserName = userName
	user.Mobile = mobile
	user.Creator = creator
	user.Editor = creator
	err = tx.Create(&user).Error
	if err != nil {
		return xerr.DbErr(xerr.UserExist)
	}
	return nil
}
