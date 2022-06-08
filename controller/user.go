package controller

import (
	"douyin/service"
	"fmt"
	"strconv"

	//"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

//注册用户
func Register(c *gin.Context) {
	//接收前端参数信息
	username := c.Query("username")
	password := c.Query("password")

	//判断用户名是否重复
	if service.CheckName(username) {
		service.ToResponse(c, service.ResponseERR("用户已经存在"))
	} else {
		//调用service注册用户接口，发送登录请求。
		if user, err := service.RegisterUser(username, password); err != nil {
			service.ToResponse(c, service.ResponseERR("用户创建失败"))
		} else {
			service.ToLoginResponse(c, service.ResponseOK("用户创建成功"), user.Id, user.Token)
		}
	}

}

func Login(c *gin.Context) {
	//接收前端参数信息
	username := c.Query("username")
	password := c.Query("password")
	// user := service.User{
	// 	UserName:     c.Query("username"),
	// 	UserPassword: c.Query("password"),
	// }
	//判断用户登录密码是否正确
	if user, err := service.LoginUser(username, password); err != nil {
		service.ToLoginResponse(c, service.ResponseERR("用户登录失败,账号密码错误"), user.Id, user.Token)
	} else {
		service.ToLoginResponse(c, service.ResponseOK("用户登录成功"), user.Id, user.Token)
	}

}

func UserInfo(c *gin.Context) {

	userid, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	if user, err := service.GetUser(userid); err == nil {
		service.ToUserResponse(c, service.ResponseOK("用户登录成功"), user)

	} else {
		fmt.Println(err)
		service.ToResponse(c, service.ResponseERR("用户不存在"))
	}
}
