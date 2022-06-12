package controller

import (
	"douyin/service"

	"github.com/gin-gonic/gin"
)

//Feed视频流接口，DemoVideos列表里有所有播放的视频列表
func Feed(c *gin.Context) {

	userid, _ := service.GetUserId(c.Query("token"))

	service.ToFeedResponse(c, service.Ok(""), service.GetVideoList(c, userid))

}
