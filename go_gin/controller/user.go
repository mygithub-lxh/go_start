package controller

import (
	"fmt"
	"go_gin/models"
	"go_gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req models.User
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var exist models.User
	if result := utils.Database.Where("username = ?", req.Username).First(&exist); result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户已经存在"})
		return
	}

	req.Password = utils.Md5Hash(req.Password)
	if result := utils.Database.Create(&req); result.Error != nil {
		fmt.Println(req)
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Registered"})
}
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var exist models.User
	if result := utils.Database.Where("username = ? ", user.Username).First(&exist); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	if exist.Password == utils.Md5Hash(user.Password) {
		token, err := utils.GenerateToken(exist.ID, exist.Username)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User Login error"})
		return
	}
}
