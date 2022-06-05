package controller

import (
	"douyin/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func FavoriteList(c *gin.Context) {

	userid, err := service.GetUserId(c.Query("token"))
	if err != nil {
		//发送错误请求
		fmt.Println(userid, err)
	}
	service.ToListResponse(c, service.ResponseOK(""), service.GetFavoriteList(userid))
}
