package task

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Task4() {
	// 连接到 SQLite 数据库
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 假设我们要查询用户 ID 为 1 的所有文章及其评论信息
	var user User
	err = db.Preload("Posts.Comments").First(&user, 1).Error
	if err != nil {
		log.Fatal("查询用户失败:", err)
	}

	// 输出用户发布的所有文章及评论信息
	fmt.Printf("用户: %s\n", user.Name)
	for _, post := range user.Posts {
		fmt.Printf("文章: %s\n", post.Title)
		for _, comment := range post.Comments {
			fmt.Printf("  评论: %s\n", comment.Content)
		}
	}
}
