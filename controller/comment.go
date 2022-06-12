package controller

import (
	"douyin/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

//评论操作调用接口
func CommentAction(c *gin.Context) {

	token := c.Query("token")
	commentid, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	actionType := c.Query("action_type")
	text := c.Query("comment_text")
	videoid, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	userid, err := service.GetUserId(token)
	if err != nil {
		service.ToResponse(c, service.Err("用户鉴权失败"))
		return
	}
	user, _ := service.GetUser(userid)

	service.Comment(c, commentid, actionType, text, user, userid, videoid)

}

//匹配每个视频对应的评论列表
func CommentList(c *gin.Context) {
	videoid, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	list := service.GetCommentList(videoid)
	service.ToCommentListResponse(c, list)
}
