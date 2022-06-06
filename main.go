package main

import (
	"douyin/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	initRouter(r)
	dao.InitDB()
	r.Run()
}
