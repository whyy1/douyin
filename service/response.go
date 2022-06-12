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

func Ok(statusmsg string) Response {
	response := Response{
		StatusCode: 0,
		StatusMsg:  statusmsg,
	}
	return response
}

func Err(statusmsg string) Response {
	response := Response{
		StatusCode: 1,
		StatusMsg:  statusmsg,
	}
	return response
}

func ToResponse(c *gin.Context, response Response) {
	c.JSON(http.StatusOK, response)
}

func ToListResponse(c *gin.Context, response Response, list []dao.Video) {
	feedresponse := VideoListResponse{
		Response:  response,
		VideoList: list,
	}
	c.JSON(http.StatusOK, feedresponse)
}
