package controller

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

//usersLoginInfo使用map存储用户信息，关键是演示的用户名+密码
//每次服务器启动时，都会清除用户数据
// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin=
// var usersLoginInfo = map[string]User{
// 	// "zhangleidouyin": {
// 	// 	Id:            1,
// 	// 	Name:          "zhanglei",
// 	// 	FollowCount:   10,
// 	// 	FollowerCount: 5,
// 	// 	IsFollow:      true,
// 	// },
// }

//var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User   User   `json:"user"`
	Avatar string `json:"avatar"`
}

//注册用户
func Register(c *gin.Context) {
	//db := GetDB()
	username := c.Query("username")
	password := c.Query("password")
	user := User{}

	//token由uuid生成
	token := uuid.New().String()

	//查询用户名是否已存在
	if err := db.First(&user, "user_name =?", username).Error; err == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		//用户不存在，则注册用户插入数据库
		newUser := User{Token: token, UserName: username, UserPassword: password, Name: username}
		_ = db.AutoMigrate(&newUser)
		error := db.Create(&newUser).Error
		if error != nil {
			fmt.Println("数据库插入失败")
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "User register success"},
			UserId:   newUser.Id,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	user := User{}
	username := c.Query("username")
	password := c.Query("password")

	if err := db.Debug().Where("user_name = ? AND user_password = ?", username, password).First(&user).Error; err == nil {
		token := uuid.New().String()
		db.Debug().First(&user).Updates(User{Token: token})
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "User login success"},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {

	userid := c.Query("user_id")
	user := User{}

	if err := db.First(&user, "id =?", userid).Error; err == nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0, StatusMsg: "UserInfo"},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
