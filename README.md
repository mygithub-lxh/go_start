## 项目简介
基于 Gin 和 GORM 的博客后端，支持用户认证、文章和评论的 CRUD 操作。

## 技术栈
- Go 1.20+
- Gin
- GORM + SQLite（可替换为 MySQL）
- JWT 认证

依赖
go mod init go_gin

# 安装 Gin 框架
go get github.com/gin-gonic/gin

# 安装 GORM 和 SQLite 驱动（或 MySQL）
go get gorm.io/gorm
go get gorm.io/driver/sqlite
# go get gorm.io/driver/mysql

# JWT 和加密库
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt

启动
go run main.go

