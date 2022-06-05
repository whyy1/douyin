package controller

import (
	//"douyin/dao"
	"douyin/service"
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {
	title := c.PostForm("title")

	userid, err := service.GetUserId(c.Query("token"))
	if err != nil {
		service.ToResponse(c, service.ResponseERR("用户鉴权失败"))
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		service.ToResponse(c, service.ResponseERR("文件上传失败"))
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", userid, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		service.ToResponse(c, service.ResponseERR(err.Error()))
		return
	}
	//传递参数写数据库增加video数据
	newvideo := service.Video{
		UserId:        userid,
		PlayUrl:       fmt.Sprintf("/static/%s", finalName),
		CoverUrl:      "http://y1-image.oss-cn-beijing.aliyuncs.com/image/2022/05/06/0593069c-f577-4bd2-9076-c07e3ad9904c.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Titile:        title,
	}

	err = db.Create(&newvideo).Error
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	service.ToResponse(c, service.ResponseOK(finalName+" 文件上传成功"))
}

func PublishList(c *gin.Context) {

	userid, err := service.GetUserId(c.Query("token"))
	if err != nil {
		//发送错误请求
		fmt.Println(userid, err)
	}
	service.ToListResponse(c, service.ResponseOK(""), service.GetPubilshList(userid))
}
