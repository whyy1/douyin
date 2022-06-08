package dao

import (
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

// func newUser(param *User) User {
// 	user := User{
// 		Id:            param.Id,
// 		UserName:      param.UserName,
// 		UserPassword:  param.UserPassword,
// 		Name:          param.Name,
// 		FollowCount:   param.FollowCount,
// 		FollowerCount: param.FollowerCount,
// 		IsFollow:      param.IsFollow,
// 		Token:         param.Token,
// 		Avatar:        param.Avatar,
// 	}
// 	return user
// }
func Register(username string, password string) (User, error) {
	//token由uuid生成
	token := uuid.New().String()
	user := User{
		UserName:     username,
		UserPassword: password,
		Token:        token,
		Name:         username,
	}
	//_ = db.AutoMigrate(&user)
	err := db.Debug().Create(&user).Error
	return user, err
}

func Find(username string) bool {
	user := User{}
	if err := db.First(&user, "user_name =?", username).Error; err == nil {
		return true
	} else {
		return false
	}
}

func Login(username string, userpass string) (User, error) {
	user := User{}
	err := db.Debug().Where("user_name = ? AND user_password = ?", username, userpass).First(&user).Error
	if err != nil {
		return User{}, err
	}
	token := uuid.New().String()
	db.Debug().First(&user, "id = ?", user.Id).Updates(User{Token: token})
	return user, err
}

func UserInfo(userid int64) (User, error) {
	user := User{}
	err := db.Debug().First(&user, "id =?", userid).Error
	return user, err
}

func UserId(token string) (*User, error) {
	user := User{}
	err := db.Debug().First(&user, "token =?", token).Error
	return &user, err
}
