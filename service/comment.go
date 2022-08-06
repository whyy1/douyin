package service

import (
	"net/http"

	"github.com/whyy1/douyin/dao"

	"github.com/gin-gonic/gin"
)

type CommentActionResponse struct {
	Response
	Comment dao.Comment `json:"comment,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []dao.Comment `json:"comment_list,omitempty"`
}

func Comment(c *gin.Context, commentid int64, actionType string, text string, user dao.User, userid int64, videoid int64) {
	if actionType == "1" {
		//对Comment表插入评论
		comment := dao.AddComment(text, user, userid, videoid)
		//增加对应的Videos的评论数
		dao.AddCommentCount(videoid)

		ToCommentActionResponse(c, comment)
	} else if actionType == "2" {
		//查看该评论或该视频是否是当前用户发布
		// comment := dao.GetComment(commentid)
		// video := dao.GetVideo(dao.Video{Id: videoid})
		// if comment.UserId != userid && video.UserId != userid {
		// 	ToResponse(ctx, ResponseERR("无权进行该操作"))
		// 	return
		// }
		//删除Comment表中评论
		dao.DeleteComment(commentid)
		//减少对应的Videos的评论数
		dao.DeductCommentCount(videoid)

		ToResponse(c, Ok("删除评论成功"))
	}

}

func GetCommentList(videoid int64) []dao.Comment {
	list := dao.CommentList(videoid)
	return list
}

func ToCommentActionResponse(c *gin.Context, comment dao.Comment) {
	response := CommentActionResponse{
		Response: Ok("评论成功"),
		Comment:  comment,
	}
	c.JSON(http.StatusOK, response)
}

func ToCommentListResponse(c *gin.Context, list []dao.Comment) {
	response := CommentListResponse{
		Response:    Ok("评论列表发布成功"),
		CommentList: list,
	}
	c.JSON(http.StatusOK, response)
}
