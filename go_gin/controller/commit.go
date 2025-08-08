package controller

import (
	"go_gin/models"
	"go_gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	comment := models.Comment{}
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := utils.Database.Create(&comment).Error; err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

func GetCommentList(c *gin.Context) {
	postId := c.Query("postId") // 获取 GET 请求参数 ?postId=123
	var comments []models.Comment

	if err := utils.Database.Where("post_id = ?", postId).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}
