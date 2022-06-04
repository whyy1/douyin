package controller

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish函数检查token，然后将上载文件保存到piblish目录,并把video数据写入数据库
// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	db := GetDB()
	token := c.PostForm("token")
	title := c.PostForm("title")

	user := User{}
	if err := db.First(&user, "token =?", token).Error; err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	newvideo := Video{
		UserId:        user.Id,
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
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

//	PublishList所有用户都有相同的发布视频列表
// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	db := GetDB()
	publishlist := []Video{}
	userid := c.Query("user_id")
	db.Debug().Preload("Author").Where("user_id =?", userid).Find(&publishlist)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: publishlist,
	})
}
