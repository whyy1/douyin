package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db gorm.DB

func main() {
	// 1.创建路由
	r := gin.Default()
	initRouter(r)
	InitDB(db)
	r.Run(":8080")
}
