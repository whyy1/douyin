package dao
import(
	"gorm.io/gorm"
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

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func (db *gorm.DB) createUser(param *User) {
	newUser := User{
	UserId        param.UserId
	UserName      string
	UserPassword  string
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count" gorm:"force:force"`
	FollowerCount int64  `json:"follower_count" gorm:"force:force"`
	IsFollow      bool   `json:"is_follow"`
	Token         string `json:"token"`
	Avatar        string `json:"avatar"`
	}
	_ = db.AutoMigrate(&newUser)
	error := db.Create(&newUser).Error
}
