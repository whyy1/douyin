package dao

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"

	"github.com/google/uuid"
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

func Register(user User) User {

	//token由uuid生成
	//token := uuid.New().String()
	user.UserPassword = getHAS256(user.UserPassword)
	user.Avatar = "https://picsum.photos/200"
	user.Name = user.UserName
	// user := User{
	// 	UserName:     username,
	// 	UserPassword: password,
	// 	Token:        token,
	// 	Name:         username,
	// 	Avatar:       "https://picsum.photos/200",
	// }
	if err := db.AutoMigrate(&user); err != nil {
		log.Println("username:%v ,User AutoMigrate fail,err=%v\n", user.UserName, err)
	}

	if err := db.Create(&user).Error; err != nil {
		log.Println("username:%v,User Create fail,err=%v", user.UserName, err)
	}

	return user
}

func Find(user User) {
	// user := User{}

	if err := db.First(&user, "user_name =?", user.Name).Error; err != nil {
		log.Println("username:%v,User Find fail,err=%v", user.UserName, err)
	}
	return
}

func Login(username string, userpass string) (User, error) {
	user := User{}

	err := db.Where("user_name = ? AND user_password = ?", username, userpass).First(&user).Error
	if err != nil {
		return User{}, err
	}
	token := uuid.New().String()
	db.First(&user, "id = ?", user.Id).Updates(User{Token: token})
	return user, err
}

func UserInfo(userid int64) (User, error) {
	user := User{}
	db.First(&user, "id =?", userid)
	err := db.First(&user, "id =?", userid).Error
	return user, err
}

//传入token返回
func UserId(token string) (*User, error) {
	user := User{}
	err := db.First(&user, "token =?", token).Error
	return &user, err
}

//增加被关注用户的follow_count数以及关注用户的follower_count数
func AddFollowCount(userid int64, followid int64) {
	user := User{}
	touser := User{}

	db.Debug().First(&user, "id = ?", userid).Update("follow_count", user.FollowCount+1)
	db.Debug().First(&touser, "id = ?", followid).Update("follower_count", touser.FollowerCount+1)
}

//减少被关注用户的follow_count数以及关注用户的follower_count数
func DeductFollowCount(userid int64, followid int64) {
	user := User{}
	touser := User{}

	db.Debug().First(&user, "id = ?", userid).Update("follow_count", user.FollowCount-1)
	db.Debug().First(&touser, "id = ?", followid).Update("follower_count", touser.FollowerCount-1)
}

//HAS256对密码进行加密
func getHAS256(password string) (has string) {
	w := sha256.New()
	io.WriteString(w, password)
	bw := w.Sum(nil)
	has = hex.EncodeToString(bw)
	return has
}
