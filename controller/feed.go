package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

//Feed视频流接口，DemoVideos列表里有所有播放的视频列表
func Feed(c *gin.Context) {
	token := c.Query("token")
	Videos := []Video{}
	favorites := []int64{}
	follow := []int64{}
	user := User{}

	//预加载，拿到videos表所有数据，并带上一个User
	//同时检索登录用户所点赞的video_id，改变点赞视频的点赞状态
	db.First(&user, "token =?", token)
	db.Debug().Table("favorites").Select("video_id").Where("user_id = ?", user.Id).Scan(&favorites)
	db.Debug().Table("follows").Select("follow_id").Where("follower_id = ?", user.Id).Scan(&follow)
	db.Debug().Preload("Author").Order("create_date desc").Find(&Videos)
	for i := range Videos {
		Videos[i].PlayUrl = "http://" + c.Request.Host + Videos[i].PlayUrl
		for _, j := range favorites {
			if Videos[i].Id == j {
				Videos[i].IsFavorite = true
				break
			}
		}
		for _, j := range follow {
			if Videos[i].Author.Id == j {
				Videos[i].Author.IsFollow = true
				break
			}
		}
	}
	//fmt.Print(Videos)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		NextTime:  time.Now().Unix(),
		VideoList: Videos,
	})
}
