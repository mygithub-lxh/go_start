package task

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 在 Post 创建时更新用户的文章数量
func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 更新用户的文章数量
	if err := tx.Model(&User{}).Where("id = ?", post.UserID).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// 在 Comment 删除时检查评论数量并更新文章状态
func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 查询文章的评论数量
	var post Post
	if err := tx.Model(&Post{}).Where("id = ?", comment.PostID).First(&post).Error; err != nil {
		return err
	}

	// 如果评论数量为 0，更新文章的评论状态为 "无评论"
	var commentCount int64
	tx.Model(&Comment{}).Where("post_id = ?", comment.PostID).Count(&commentCount)

	if commentCount == 0 {
		// 更新文章的评论状态
		if err := tx.Model(&post).Update("comment_count", 0).Error; err != nil {
			return err
		}
	}

	return nil
}

func Task5() {
	// 连接到 SQLite 数据库
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 自动迁移模型到数据库，创建对应的表
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatal("自动迁移失败:", err)
	}

	// 创建一个示例用户
	user := User{Name: "John Doe", Email: "john@example.com", Password: "securepassword"}
	db.Create(&user)

	// 创建文章并自动更新用户的文章数量
	post := Post{Title: "First Post", Content: "This is my first post!", UserID: user.ID}
	db.Create(&post)

	// 创建评论
	comment := Comment{Content: "Great post!", PostID: post.ID}
	db.Create(&comment)

	// 删除评论并检查评论数量
	db.Delete(&comment)

	// 输出用户的文章数量和文章的评论数量
	var updatedUser User
	db.First(&updatedUser, user.ID)
	fmt.Printf("用户 %s 发布的文章数量: %d\n", updatedUser.Name, updatedUser.PostCount)

	var updatedPost Post
	db.First(&updatedPost, post.ID)
	fmt.Printf("文章 '%s' 的评论数量: %d\n", updatedPost.Title, updatedPost.CommentCount)
}
