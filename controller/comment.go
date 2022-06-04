package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

//
func CommentAction(c *gin.Context) {

	token := c.Query("token")
	actionType := c.Query("action_type")
	commentid := c.Query("comment_id")
	video_id, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	user := User{}

	if err := db.First(&user, "token =?", token).Error; err == nil {
		if actionType == "1" {
			//对Comment表插入评论
			text := c.Query("comment_text")
			comment := Comment{
				User:       user,
				Content:    text,
				UserId:     user.Id,
				VideoId:    video_id,
				CreateDate: time.Now().Format("2006-01-02 15:04:05"),
			}
			if err := db.Create(&comment).Error; err != nil {
				fmt.Println("评论插入失败", err)
			}
			//增加对应的Videos的评论数
			videos := Video{}
			db.Debug().First(&videos, "id=?", video_id).Update("comment_count", videos.CommentCount+1)

			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
				Comment: comment})
		} else if actionType == "2" {
			//减少comments表中的评论
			db.Debug().Where("id =?", commentid).Delete(&Comment{})
			//减少对应的Videos的评论数
			videos := Video{}
			db.Debug().First(&videos, "id=?", video_id).Update("comment_count", videos.CommentCount-1)
			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0}})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

//匹配每个视频对应的评论列表
func CommentList(c *gin.Context) {
	List := []Comment{}
	VideorId := c.Query("video_id")
	db.Debug().Preload("User").Order("create_date desc").Where("video_id=?", VideorId).Find(&List)
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: List,
	})
}
