package service

import (
	"douyin/dao"
	"net/http"

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

func Comment(ctx *gin.Context, commentid int64, actionType string, text string, user dao.User, userid int64, videoid int64) {
	if actionType == "1" {
		//对Comment表插入评论
		comment, _ := dao.AddComment(text, user, userid, videoid)
		//增加对应的Videos的评论数
		dao.AddCommentCount(videoid)

		ToCommentActionResponse(ctx, comment)
	} else if actionType == "2" {
		//删除Comment表中评论
		dao.DeleteComment(commentid)
		//减少对应的Videos的评论数
		dao.DeductCommentCount(videoid)

		ToResponse(ctx, ResponseOK("删除评论成功"))
	}

}

func GetCommentList(videoid int64) []dao.Comment {
	list, _ := dao.CommentList(videoid)
	return list
}

func ToCommentActionResponse(ctx *gin.Context, comment dao.Comment) {
	response := CommentActionResponse{
		Response: ResponseOK("评论成功"),
		Comment:  comment,
	}
	ctx.JSON(http.StatusOK, response)
}

func ToCommentListResponse(ctx *gin.Context, list []dao.Comment) {
	response := CommentListResponse{
		Response:    ResponseOK("评论列表发布成功"),
		CommentList: list,
	}
	ctx.JSON(http.StatusOK, response)
}
