package controller

import (
	"douyin/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

//Feed视频流接口，DemoVideos列表里有所有播放的视频列表
func Feed(c *gin.Context) {

	userid, err := service.GetUserId(c.Query("token"))
	if err != nil {
		//发送错误请求
		fmt.Println(userid, err)
	}

	service.ToFeedResponse(c, service.ResponseOK(""), service.GetVideoList(c, userid))
}
