package db

import (
	"context"
	"log"
	"server/common/config"
	"testing"
)

func TestNewModel(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	m := NewModel()
	if m == nil {
		t.Error("NewModel error")
	}
	userInfo, err := m.FindUserByUsername(context.Background(), "test")
	if err != nil {
		t.Error(err)
		return
	}
	log.Print(userInfo)
}
