package controller

import (
	"douyin/service"

	"github.com/gin-gonic/gin"
)

//Feed视频流接口，DemoVideos列表里有所有播放的视频列表
func Feed(c *gin.Context) {

	userid, _ := service.GetUserId(c.Query("token"))
	service.ToFeedResponse(c, service.ResponseOK(""), service.GetVideoList(c, userid))
	// c.JSON(http.StatusOK, service.FeedResponse{
	// 	Response:  service.Response{StatusCode: 0},
	// 	NextTime:  time.Now().Unix(),
	// 	VideoList: service.GetVideoList(c, userid),
	// })
}
