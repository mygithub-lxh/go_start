package main

import (
	"go_gin/controller"
	"go_gin/middleware"
	"go_gin/utils"

	"github.com/gin-gonic/gin"
)

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo3LCJ1c2VybmFtZSI6Imx4aDciLCJleHAiOjE3NTQ2NzM3NjQsIm5iZiI6MTc1NDU4NzM2NCwiaWF0IjoxNzU0NTg3MzY0fQ.-O5BGP0OaSwRdRQAGxet503IpjDb0br7X-cGzYKILJw
func main() {
	utils.InitDB()
	router := gin.Default()
	router.Use(middleware.ErrorHandlingMiddleware())

	api := router.Group("/api")
	{
		api.POST("/login", controller.Login)
		api.POST("/register", controller.Register)
		api.GET("/post/getList", controller.GetPostList)
		api.GET("/post/getContent", controller.GetPostContent)
		api.GET("/comment/getList", controller.GetCommentList)
	}
	authApi := router.Group("/api", middleware.JWTAuthMiddleware())
	{
		authApi.PUT("/post/create", controller.CreatePost)
		authApi.POST("/post/update", controller.UpdatePost)
		authApi.DELETE("/post/delete", controller.DeletePost)
		authApi.PUT("/comment/create", controller.CreateComment)
	}
	router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
