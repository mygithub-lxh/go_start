package controller

import (
	"fmt"
	"go_gin/models"
	"go_gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPostList(c *gin.Context) {
	var posts models.Post

	if result := utils.Database.Model(&models.Post{}).Select("title", "content").Find(&posts); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": posts})
}
func GetPostContent(c *gin.Context) {
	var post = models.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or missing id"})
		return
	}

	if result := utils.Database.Select(&models.Post{}, post.ID); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}
func CreatePost(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := utils.Database.Create(&post).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
	return

}

func UpdatePost(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := utils.Database.Where("id = ?", post.ID).Updates(post).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
	return

}

func DeletePost(c *gin.Context) {
	var postId struct {
		ID uint `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&postId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or missing id"})
		return
	}

	fmt.Println(postId.ID)

	if err := utils.Database.Delete(&models.Post{}, postId.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
