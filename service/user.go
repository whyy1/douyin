package service

import (
	"errors"
	"net/http"

	"github.com/whyy1/douyin/dao"
	"github.com/whyy1/douyin/jwt"

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

// func CheckName(parameter User) bool {
// 	user := dao.User{
// 		Name: parameter.UserName,
// 	}
// 	if dao.Find(user) == true {
// 		return true
// 	}
// 	return false
// }

//传入账号密码
func RegisterUser(parameter User) (dao.User, error) {
	user := dao.User{
		UserName:     parameter.UserName,
		Name:         parameter.UserName,
		UserPassword: parameter.UserPassword,
	}
	if dao.FindName(user) == true {
		err := errors.New("用户名已经存在")
		return user, err
	}
	user, err := dao.Register(user)
	return user, err
}

func LoginUser(parameter User) (dao.User, error) {
	//查询用户名密码是否正确
	user := dao.User{
		Name:         parameter.UserName,
		UserName:     parameter.UserName,
		UserPassword: parameter.UserPassword,
	}
	user, err := dao.Login(user)

	return user, err
}

//传入用户ID返回用户信息
func GetUser(userid int64) (dao.User, error) {
	user, err := dao.UserInfo(userid)
	return user, err
}

//传入token返回用户id
func GetUserId(token string) (int64, error) {

	//var r_token string

	//判断token是否过期,过期直接返回
	id, pd := jwt.QueryToken(token)
	if !pd {
		err := errors.New("token已失效")
		return id, err
	}
	//Redis不存在token,则查询Mysql中是否存在token,找到则返回
	if str, _ := dao.GetToken(id); str == token {
		return id, nil
	}
	//Redis不存在token,则查询Mysql中是否存在token,找到则返回
	if user, _ := dao.UserInfo(id); user.Token == token {
		dao.SetToken(id, user.Token)
		return id, nil
	}

	//user, err := dao.UserId(token)
	//return user.Id, err
	//都没找到token,失效
	return id, errors.New("token已失效")
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
