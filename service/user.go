package service

import (
	"douyin/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserId        int64 `json:"id,omitempty"`
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

func RegisterUser(param *User) (*dao.User, error) {

	//var user User
	user, err := dao.Register(&dao.User{
		UserName:     param.UserName,
		UserPassword: param.UserPassword,
		Name:         param.Name,
	})

	//查询用户名是否已存在
	if user, err := dao.Find(user); err == nil {
		return user, err
	}
	if user, err := dao.Register(user); err != nil {
		return user, err
	}
	return user, err
}

func LoginUser(param *User) (*dao.User, error) {
	user, err := dao.Login(&dao.User{
		UserName:     param.UserName,
		UserPassword: param.UserPassword,
	})
	//查询用户名密码是否正确
	///_, err := dao.Login(&param)
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
	return user.UserId, err
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
