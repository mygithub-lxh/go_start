package utils

import (
	"go_gin/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模型
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	Database = db
}
