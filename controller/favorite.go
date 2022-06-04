package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//判断登录用户，对点赞/取消点赞做出相应操作
func FavoriteAction(c *gin.Context) {
	db := GetDB()
	token := c.Query("token")
	videoid := c.Query("video_id")
	actiontype := c.Query("action_type")
	user := User{}
	videos := Video{}

	if err := db.First(&user, "token =?", token).Error; err == nil {
		if actiontype == "1" {
			//增加Favorite表中记录
			favorite := Favorite{UserId: user.Id, VideoId: videoid}
			if err := db.Create(&favorite).Error; err != nil {
				fmt.Println("点赞失败", err)
			}
			//根据video_id,增加对应的Video的点赞数
			db.Debug().First(&videos, "id = ?", videoid).Updates(Video{FavoriteCount: videos.FavoriteCount + 1})
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		} else if actiontype == "2" {
			//删除Favorite表中记录
			if err := db.Debug().Where("user_id = ?", user.Id).Where("video_id = ?", videoid).Delete(&Favorite{}).Error; err != nil {
				fmt.Println("点赞失败", err)
			}
			//根据video_id,减少对应的Video的点赞数
			db.Debug().First(&videos, "id=?", videoid).Updates(Video{FavoriteCount: videos.FavoriteCount - 1})
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// 拿到登录用户id，检索favorites表中点赞视频的id，然后返回视频列表
func FavoriteList(c *gin.Context) {
	userid := c.Query("user_id")
	favoritelist := []Video{}
	var arr = []int64{}

	db.Debug().Table("favorites").Select("video_id").Where("user_id = ?", userid).Scan(&arr)
	db.Debug().Preload("Author").Where("id IN ?", arr).Find(&favoritelist)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0, StatusMsg: "FavoriteList get success",
		},
		VideoList: favoritelist,
	})
}
