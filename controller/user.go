package controller

import (
	"douyin/service"
	"strconv"

	//"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

//注册用户
func Register(c *gin.Context) {
	//接收参数信息
	username := c.Query("username")
	password := c.Query("password")

	//判断用户名是否存在
	if service.CheckName(username) {
		service.ToResponse(c, service.Err("用户已经存在"))
	} else {
		//用户名不存在则新建用户
		if user, err := service.RegisterUser(username, password); err != nil {
			service.ToResponse(c, service.Err("用户创建失败"))
		} else {
			service.ToLoginResponse(c, service.Ok("用户创建成功"), user.Id, user.Token)
		}
	}
}

func Login(c *gin.Context) {
	//接收参数信息
	username := c.Query("username")
	password := c.Query("password")

	//判断用户登录密码是否正确
	user, err := service.LoginUser(username, password)
	if err != nil {
		service.ToLoginResponse(c, service.Err("用户登录失败,账号密码错误"), user.Id, user.Token)
		return
	}

	service.ToLoginResponse(c, service.Ok("用户登录成功"), user.Id, user.Token)
}

func UserInfo(c *gin.Context) {

	userid, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	if user, err := service.GetUser(userid); err == nil {
		service.ToUserResponse(c, service.Ok("用户登录成功"), user)
	} else {
		service.ToResponse(c, service.Err("用户不存在"))
	}
}
