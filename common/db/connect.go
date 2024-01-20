package db

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"server/common/config"
	"server/ent"
	"server/ent/migrate"
)

type customModel struct {
	client *ent.Client
}

func NewModel() Model {
	dsn := "host=%s user=%s password=%s dbname=%s port=%d  TimeZone=Asia/Shanghai sslmode=disable"
	c := config.RunData.DbConfig
	dsn = fmt.Sprintf(dsn, c.Addr, c.Username, c.Password, c.DbName, c.Port)
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Panic(err)
	}
	err = CreateTable(client)
	if err != nil {
		log.Panic(err)
	}
	return &customModel{
		client: client,
	}
}

func CreateTable(client *ent.Client) error {
	err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
		migrate.WithDropIndex(true),
	)
	return err
}
