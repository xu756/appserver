package db

import (
	"context"
	"server/common/config"
	"testing"
)

func TestCustomModel_FindUserByMobile(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	model := TestModel()
	user, err := model.FindUserByMobile(context.Background(), "17337687416")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user)
}

func TestCustomModel_CreateUser(t *testing.T) {
	model := TestModel()
	err := model.CreateUser(context.Background(), "xu756", "17337687416", 111)
	if err != nil {
		t.Error(err)
		return
	}

}
