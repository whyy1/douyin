package dao

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"

	"github.com/whyy1/douyin/jwt"
	"gorm.io/gorm"
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

func Register(user User) (User, error) {

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
		log.Println("username:", user.UserName, ",User AutoMigrate fail,err=", err)
		return user, err
	}

	if err := db.Create(&user).Error; err != nil {
		log.Println("username:", user.UserName, ",User Create fail,err=", err)
		return user, err
	}

	return user, nil
}

func FindName(user User) bool {
	// user := User{}

	if err := db.First(&user, "user_name =?", user.Name).Error; err != nil {
		//log.Println("username:%v,User Find fail,err=%v", user.UserName, err)
		return false
	}
	return true
}

func Login(user User) (User, error) {
	//user := User{}
	user.UserPassword = getHAS256(user.UserPassword)
	err := db.Where("user_name = ? AND user_password = ?", user.Name, user.UserPassword).First(&user).Error
	if err != nil {
		return User{}, err
	}
	token := jwt.GetToken(user.Id)
	SetToken(user.Id, user.Token)
	db.First(&user, "id = ?", user.Id).Updates(User{Token: token})
	user.Token = token
	return user, err
}

func UserInfo(userid int64) (User, error) {
	user := User{
		Id: userid,
	}
	//	db.First(&user, "id =?", userid)
	if err := db.First(&user, "id =?", user.Id).Error; err != nil {
		return user, err
	}
	return user, nil
}

//传入token返回
// func UserId(token string) (*User, error) {
// 	user := User{}
// 	err := db.First(&user, "token =?", token).Error
// 	return &user, err
// }

//判断Token有没有找到
// func FindToken(id int64) (int64, error) {
// 	user,_ := UserInfo(id)
// 	err := db.First(&user, "id =?", token).Error
// 	return user.Id, err
// }

//增加被关注用户的follow_count数以及关注用户的follower_count数
func AddFollowCount(userid int64, followid int64) {
	user := User{}
	touser := User{}

	if err := db.Transaction(func(tx *gorm.DB) error {

		if err := tx.First(&user, "id = ?", userid).Update("follow_count", user.FollowCount+1).Error; err != nil {
			log.Println("用户ID:", user.Id, ",关注操作-用户错误，err:", err)
			return err
		}

		if err := tx.First(&touser, "id = ?", followid).Update("follower_count", touser.FollowerCount+1).Error; err != nil {
			log.Println("被关注的用户ID:", touser.Id, ",关注操作-被关注用户错误，err=", err)
			return err
		}
		return nil
	}); err != nil {
		log.Println("用户ID:", user.Id, " 被关注的用户ID", touser.Id, "，关注操作错误 err=", err)
	}
	// db.Debug().First(&user, "id = ?", userid).Update("follow_count", user.FollowCount+1)
	// db.Debug().First(&touser, "id = ?", followid).Update("follower_count", touser.FollowerCount+1)
}

//减少被关注用户的follow_count数以及关注用户的follower_count数
func DeductFollowCount(userid int64, followid int64) {
	user := User{}
	touser := User{}

	if err := db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.First(&user, "id = ?", userid).Update("follow_count", user.FollowCount-1).Error; err != nil {
			// 返回任何错误都会回滚事务
			log.Println("用户ID:", user.Id, ",取消关注操作-用户错误，err:", err)
			return err
		}

		if err := tx.First(&touser, "id = ?", followid).Update("follower_count", touser.FollowerCount-1).Error; err != nil {
			log.Println("被关注的用户ID:", touser.Id, ",取消关注操作-被关注用户错误，err=", err)
			return err
		}
		// 返回 nil 提交事务
		return nil
	}); err != nil {
		log.Println("用户ID:", user.Id, " 被关注的用户ID", touser.Id, "，取消关注操作错误 err=", err)
	}
	// db.Debug().First(&user, "id = ?", userid).Update("follow_count", user.FollowCount-1)
	// db.Debug().First(&touser, "id = ?", followid).Update("follower_count", touser.FollowerCount-1)
}

//HAS256对密码进行加密
func getHAS256(password string) (has string) {
	w := sha256.New()
	_, _ = io.WriteString(w, password)
	bw := w.Sum(nil)
	has = hex.EncodeToString(bw)
	return has
}
