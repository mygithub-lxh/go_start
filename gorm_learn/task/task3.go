package task

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 模型，增加了 CommentCount 字段
type User struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `gorm:"size:100;not null"`
	Email        string    `gorm:"size:100;unique;not null"`
	Password     string    `gorm:"size:100;not null"`
	PostCount    int       `gorm:"default:0"` // 用户的文章数量
	CommentCount int       `gorm:"default:0"` // 用户的评论数量
	Posts        []Post    // 用户与文章的一对多关系
	Comments     []Comment // 用户与评论的一对多关系
}

// Post 模型
type Post struct {
	ID           uint      `gorm:"primaryKey"`
	Title        string    `gorm:"size:255;not null"`
	Content      string    `gorm:"type:text;not null"`
	UserID       uint      // 外键：UserID 表示这篇文章属于哪个用户
	CommentCount int       `gorm:"default:0"` // 用户的评论数量
	User         User      // 关联到 User
	Comments     []Comment // 文章与评论的一对多关系
}

// Comment 模型
type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"type:text;not null"`
	PostID  uint   // 外键：PostID 表示这条评论属于哪篇文章
	Post    Post   // 关联到 Post
	UserID  uint   // 外键：UserID 表示这条评论属于哪个用户
	User    User   // 关联到 User
}

func Task3() {
	// 连接到 SQLite 数据库（如果没有该数据库，GORM 会自动创建）
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 自动迁移模型到数据库，创建对应的表
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatal("自动迁移失败:", err)
	}

	fmt.Println("数据库表创建成功！")
}
