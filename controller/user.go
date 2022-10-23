package controller

import (
	"strconv"

	"github.com/whyy1/douyin/service"

	//"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

//注册用户
func Register(c *gin.Context) {
	// 接收参数信息
	username := c.Query("username")
	password := c.Query("password")
	user := service.User{
		UserName:     username,
		UserPassword: password,
		Name:         username,
	}
	//判断用户名是否存在
	// if service.CheckName(user.Name) {
	// 	service.ToResponse(c, service.Err("用户已经存在"))
	// 	return
	// }

	//用户名不存在则新建用户
	if user, err := service.RegisterUser(user); err != nil {
		service.ToResponse(c, service.Err("用户创建失败,用户名已存在"))
	} else {
		service.ToLoginResponse(c, service.Ok("用户创建成功"), user.Id, user.Token)
	}
}

func Login(c *gin.Context) {
	//接收参数信息
	username := c.Query("username")
	password := c.Query("password")
	user := service.User{
		UserName:     username,
		UserPassword: password,
		Name:         username,
	}

	//判断用户登录密码是否正确
	if user, err := service.LoginUser(user); err != nil {
		service.ToLoginResponse(c, service.Err("用户登录失败,账号密码错误"), user.Id, user.Token)
		return
	} else {
		service.ToLoginResponse(c, service.Ok("用户登录成功"), user.Id, user.Token)
	}

}

func UserInfo(c *gin.Context) {

	userid, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	if user, err := service.GetUser(userid); err != nil {
		service.ToResponse(c, service.Err("用户不存在"))
	} else {
		service.ToUserResponse(c, service.Ok("用户登录成功"), user)
	}
}
