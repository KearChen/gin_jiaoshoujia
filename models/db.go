package models

import (
	"fmt"
	"gin_jiaoshoujia/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dbType := config.Config.GetString("database.type")

	switch dbType {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Config.GetString("database.mysql.user"),
			config.Config.GetString("database.mysql.password"),
			config.Config.GetString("database.mysql.host"),
			config.Config.GetInt("database.mysql.port"),
			config.Config.GetString("database.mysql.dbname"),
		)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "sqlite":
		dbPath := config.Config.GetString("database.sqlite.path")
		DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	default:
		log.Fatalf("不支持的数据库类型: %s", dbType)
	}

	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	log.Println("数据库连接成功")
	// 自动迁移
	AutoMigrate()
}
