package db

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"server/common/config"
	"server/common/model"
	"time"
)

type customModel struct {
	Db *gorm.DB
}

func NewModel() Model {
	dsn := "host=%s user=%s password=%s dbname=%s port=%d  TimeZone=Asia/Shanghai"
	c := config.RunData.DbConfig
	dsn = fmt.Sprintf(dsn, c.Addr, c.Username, c.Password, c.DbName, c.Port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		log.Panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDb.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDb.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDb.SetConnMaxLifetime(time.Second)
	log.Print("【 数据库连接成功 】")
	err = CreateTable(db)
	if err != nil {
		panic(err)
	}
	return &customModel{
		Db: db,
	}
}

func CreateTable(db *gorm.DB) error {
	err := db.AutoMigrate(&model.UserModel{})
	if err != nil {
		klog.Debugf("【 创建表失败 %s 】 ", "user")
		return err
	}
	err = db.AutoMigrate(&model.RoleModel{})
	if err != nil {
		klog.Debugf("【 创建表失败 %s 】 ", "role")
		return err
	}
	err = db.AutoMigrate(&model.UserRoleModel{})
	if err != nil {
		klog.Debugf("【 创建表失败 %s 】 ", "user_role")
		return err
	}
	return nil
}

func TestModel() Model {
	dsn := "host=%s user=%s password=%s dbname=%s port=%d  TimeZone=Asia/Shanghai"
	c := config.RunData.DbConfig
	dsn = fmt.Sprintf(dsn, c.Addr, c.Username, c.Password, c.DbName, c.Port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		log.Panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDb.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDb.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDb.SetConnMaxLifetime(time.Second)
	log.Print("【 数据库连接成功 】")
	err = CreateTable(db)
	if err != nil {
		panic(err)
	}
	return &customModel{
		Db: db,
	}
}
