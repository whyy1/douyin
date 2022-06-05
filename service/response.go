package service

import (
	"douyin/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func ResponseOK(statusmsg string) Response {
	response := Response{
		StatusCode: 0,
		StatusMsg:  statusmsg,
	}
	return response
}
func ResponseERR(statusmsg string) Response {
	response := Response{
		StatusCode: 1,
		StatusMsg:  statusmsg,
	}
	return response
}

func ToResponse(ctx *gin.Context, response Response) {
	ctx.JSON(http.StatusOK, response)
}

func ToListResponse(ctx *gin.Context, response Response, list []dao.Video) {
	feedresponse := VideoListResponse{
		Response:  response,
		VideoList: list,
	}
	ctx.JSON(http.StatusOK, feedresponse)
}
