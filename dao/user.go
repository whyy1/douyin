package dao

import (
	"github.com/google/uuid"
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

func newUser(param *User) User {
	user := User{
		UserId:        param.UserId,
		UserName:      param.UserName,
		UserPassword:  param.UserPassword,
		Name:          param.Name,
		FollowCount:   param.FollowCount,
		FollowerCount: param.FollowerCount,
		IsFollow:      param.IsFollow,
		Token:         param.Token,
		Avatar:        param.Avatar,
	}
	return user
}
func Register(param *User) (*User, error) {
	//token由uuid生成
	token := uuid.New().String()
	user := newUser(param)
	user.Token = token
	//_ = db.AutoMigrate(&newUser)
	err := db.Debug().Create(&user).Error
	return &user, err
}

func Find(param *User) (*User, error) {
	user := newUser(param)
	err := db.Debug().First(&user, "user_name =?", user.UserName).Error
	return &user, err
}
func Login(param *User) (*User, error) {
	user := newUser(param)
	err := db.Debug().Where("user_name = ? AND user_password = ?", user.UserName, user.UserPassword).First(&user).Error
	token := uuid.New().String()
	db.Debug().First(&user, "user_id = ?", user.UserId).Updates(User{Token: token})
	return &user, err
}
func UserInfo(param *User) (*User, error) {
	user := newUser(param)
	err := db.Debug().First(&user, "user_id =?", user.UserId).Error
	return &user, err
}
func UserId(param *User) (*User, error) {
	user := newUser(param)
	err := db.Debug().First(&user, "token =?", user.Token).Error
	return &user, err
}
