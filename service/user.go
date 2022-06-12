package service

import (
	"douyin/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id            int64 `json:"id,omitempty"`
	UserName      string
	UserPassword  string
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count" gorm:"force:force"`
	FollowerCount int64  `json:"follower_count" gorm:"force:force"`
	IsFollow      bool   `json:"is_follow"`
	Token         string `json:"token"`
	Avatar        string `json:"avatar"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}
type UserResponse struct {
	Response
	dao.User `json:"user"`
}

func CheckName(username string) bool {
	if err := dao.Find(username); err == nil {
		return true
	}
	return false
}

//传入账号密码
func RegisterUser(username string, password string) (dao.User, error) {
	user, err := dao.Register(username, password)
	return user, err
}

func LoginUser(username string, userpass string) (dao.User, error) {
	//查询用户名密码是否正确
	user, err := dao.Login(username, userpass)

	return user, err
}

//传入用户ID返回用户信息
func GetUser(userid int64) (dao.User, error) {
	user, err := dao.UserInfo(userid)
	return user, err
}

//传入token返回用户id
func GetUserId(token string) (int64, error) {
	user, err := dao.UserId(token)
	return user.Id, err
}

func ToLoginResponse(ctx *gin.Context, response Response, user_id int64, token string) {
	loginresponse := UserLoginResponse{
		Response: response,
		UserId:   user_id,
		Token:    token,
	}
	ctx.JSON(http.StatusOK, loginresponse)
}
func ToUserResponse(ctx *gin.Context, response Response, user dao.User) {
	loginresponse := UserResponse{
		Response: response,
		User:     user,
	}
	ctx.JSON(http.StatusOK, loginresponse)
}
