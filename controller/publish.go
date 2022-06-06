package controller

import (
	//"douyin/dao"

	"douyin/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {
	title := c.PostForm("title")

	userid, err := service.GetUserId(c.PostForm("token"))
	if err != nil {
		service.ToResponse(c, service.ResponseERR("用户鉴权失败"))
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		service.ToResponse(c, service.ResponseERR("视频接收失败"))
		return
	}

	service.PublishVideo(data, userid)

	//传递参数写数据库增加video数据
	playurl := fmt.Sprintf("/static/%v", 1)
	if _, err := service.SaveVideo(userid, playurl, title); err != nil {
		panic("视频存入数据库失败" + err.Error())
	}
	service.ToResponse(c, service.ResponseOK(" 视频上传成功"))
}

func PublishList(c *gin.Context) {

	userid, err := service.GetUserId(c.Query("token"))
	if err != nil {
		//发送错误请求
		fmt.Println(userid, err)
	}
	service.ToListResponse(c, service.ResponseOK(""), service.GetPubilshList(userid))
}
