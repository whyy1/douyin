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
	user := service.User{
		UserName:     c.Query("username"),
		UserPassword: c.Query("password"),
	}
	//调用service注册用户接口，如果成功发送登录请求。
	if user, err := service.RegisterUser(&user); err != nil {
		fmt.Println(err)
		service.ToLoginResponse(c, service.ResponseERR("用户已经存在"), user.UserId, user.Token)
	}
	service.ToLoginResponse(c, service.ResponseOK("用户创建成功"), user.UserId, user.Token)
}

func Login(c *gin.Context) {
	//接收前端参数信息
	// username := c.Query("username")
	// password := c.Query("password")
	user := service.User{
		UserName:     c.Query("username"),
		UserPassword: c.Query("password"),
	}
	//先判断用户登录密码是否正确
	if user, err := service.LoginUser(&user); err == nil {
		service.ToLoginResponse(c, service.ResponseOK("用户登录成功"), user.UserId, user.Token)

	} else {
		fmt.Println(err)
		service.ToLoginResponse(c, service.ResponseERR("用户登录失败"), user.UserId, user.Token)
	}
}

func UserInfo(c *gin.Context) {

	userid, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	user := service.User{
		UserId: userid,
	}

	if user, err := service.GetUser(&user); err == nil {
		service.ToUserResponse(c, service.ResponseOK("用户登录成功"), user)

	} else {
		fmt.Println(err)
		service.ToLoginResponse(c, service.ResponseERR("用户不存在"), 0, "")
	}
}
