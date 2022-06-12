package service

import (
	"douyin/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []dao.User `json:"user_list"`
}

func Relation(actiontype string, userid int64, followid int64) {
	if actiontype == "1" {
		dao.AddFollow(userid, followid)
		dao.AddFollowCount(userid, followid)
	} else if actiontype == "2" {
		dao.DeleteFollow(userid, followid)
		dao.DeductFollowCount(userid, followid)
	}
}
func GetFollowList(userid int64, authorid int64) []dao.User {

	followlist := dao.FollowList(userid, authorid)

	return followlist
}

func GetFollowerList(userid int64, authorid int64) []dao.User {

	followerlist := dao.FollowerList(userid, authorid)

	return followerlist
}

func ToUserListResponse(ctx *gin.Context, list []dao.User) {
	response := UserListResponse{
		Response: Ok(""),
		UserList: list,
	}
	ctx.JSON(http.StatusOK, response)
}
