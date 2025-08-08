package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered1: %v\n", err)
				debug.PrintStack()
				c.JSON(http.StatusNotImplemented, gin.H{
					"code":    500,
					"message": "服务器内部错误",
				})
				c.Abort()
			}
		}()

		// 继续后续中间件或处理器
		c.Next()
	}
}
