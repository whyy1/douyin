package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/whyy1/douyin/service"
)

//判断登录用户，对点赞/取消点赞做出相应操作
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoid, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actiontype := c.Query("action_type")

	userid, err := service.GetUserId(token)
	if err != nil {
		service.ToResponse(c, service.Err("token过期或失效"))
		return
	}
	service.FavoriteAction(actiontype, userid, videoid)
	service.ToResponse(c, service.Ok("点赞成功"))
}
func FavoriteList(c *gin.Context) {

	userid, err := service.GetUserId(c.Query("token"))
	if err != nil {
		//发送错误请求
		fmt.Println(userid, err)
	}
	service.ToListResponse(c, service.Ok(""), service.GetFavoriteList(userid))
}
