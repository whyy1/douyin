package service

import (
	"douyin/dao"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []dao.Video `json:"video_list,omitempty"`
	NextTime  int64       `json:"next_time,omitempty"`
}

func GetVideoList(c *gin.Context, userid int64) []dao.Video {
	videos := dao.GetVideoList()
	favorites := dao.VideoFavorite(userid)
	follow := dao.Videofollow(userid)
	for i := range videos {
		videos[i].PlayUrl = "http://" + c.Request.Host + videos[i].PlayUrl
		for _, j := range favorites {
			if videos[i].UserId == j {
				videos[i].IsFavorite = true
				break
			}
		}
		for _, j := range follow {
			if videos[i].Author.UserId == j {
				videos[i].Author.IsFollow = true
				break
			}
		}
	}
	return videos
}
func ToFeedResponse(ctx *gin.Context, response Response, list []dao.Video) {
	feedresponse := FeedResponse{
		Response:  response,
		VideoList: list,
		NextTime:  time.Now().Unix(),
	}
	ctx.JSON(http.StatusOK, feedresponse)
}
