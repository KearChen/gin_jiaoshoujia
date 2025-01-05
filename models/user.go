package models

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);not null" json:"name"`
	Email string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Age   int    `json:"age"`
}

// AutoMigrate 自动迁移
func AutoMigrate() {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		panic("数据库迁移失败: " + err.Error())
	}
}
