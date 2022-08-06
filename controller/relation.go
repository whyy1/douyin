package controller

import (
	"strconv"

	"github.com/whyy1/douyin/service"

	"github.com/gin-gonic/gin"
)

func RelationAction(c *gin.Context) {
	token := c.Query("token")
	followid, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actiontype := c.Query("action_type")

	userid, err := service.GetUserId(token)
	if err != nil {
		service.ToResponse(c, service.Err("用户鉴权失败"))
		return
	}
	service.Relation(actiontype, userid, followid)
	if actiontype == "1" {
		service.ToResponse(c, service.Err("关注成功"))
	} else if actiontype == "2" {
		service.ToResponse(c, service.Err("取消关注成功"))
	}
}

func FollowList(c *gin.Context) {
	//通过token拿到当前登录用户id
	//查看作者id为user_id
	token := c.Query("token")
	authorid, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	userid, err := service.GetUserId(token)
	if err != nil {
		service.ToResponse(c, service.Err("用户鉴权失败"))
		return
	}

	followlist := service.GetFollowList(userid, authorid)
	service.ToUserListResponse(c, followlist)
}
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	authorid, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	userid, err := service.GetUserId(token)
	if err != nil {
		service.ToResponse(c, service.Err("用户鉴权失败"))
		return
	}

	followerlist := service.GetFollowerList(userid, authorid)
	service.ToUserListResponse(c, followerlist)
}
