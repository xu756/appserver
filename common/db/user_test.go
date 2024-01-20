package db

import (
	"context"
	"server/common/config"
	"server/ent"
	"testing"
)

func TestCustomModel_CreateUser(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	m := NewModel()
	if m == nil {
		t.Error("NewModel error")
	}
	user, err := m.CreateUser(context.TODO(), "test", "test", "12345678901", 0)
	if err != nil {
		t.Error(err)
		t.Error(ent.IsValidationError(err))
		return
	}
	t.Log(user)
}
