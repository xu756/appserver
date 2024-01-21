package db

import (
	"context"
	"server/common/config"
	"server/ent"
	"testing"
)

func TestCustomModel_FindUserByUsername(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	m := NewModel()
	userInfo, err := m.FindUserByUsername(context.Background(), "test")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(userInfo)
}

func TestCustomModel_FindUserByMobile(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	m := NewModel()
	userInfo, err := m.FindUserByMobile(context.Background(), "12345678101")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(userInfo)
}
func TestCustomModel_CreateUser(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	m := NewModel()
	user, err := m.CreateUser(context.TODO(), "test", "test", "12345678101", 0)
	if err != nil {
		t.Error(err)
		t.Error(ent.IsConstraintError(err))
		return
	}
	t.Log(user)
}
