package controller

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64    `json:"id,omitempty"`
	Author        User     `json:"author" gorm:"foreignkey:UserId"`
	PlayUrl       string   `json:"play_url,omitempty"`
	CoverUrl      string   `json:"cover_url,omitempty"`
	FavoriteCount int64    `json:"favorite_count" gorm:"force:force"`
	CommentCount  int64    `json:"comment_count" gorm:"force:force"`
	Favorite      Favorite ` gorm:"foreignkey:UserId"`
	IsFavorite    bool     `json:"is_favorite"`
	UserId        int64    `gorm:"not null"`
	Titile        string   `json:"title,omitempty"`
	CreateDate    int64    `gorm:"autoCreateTime"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user" gorm:"foreignkey:UserId"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	UserId     int64  `gorm:"not null"`
	VideoId    int64  //`gorm:"not null" gorm:"foreignkey:VideorId"`
}

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
type Favorite struct {
	UserId  int64  `gorm:"not null"`
	VideoId string `gorm:"not null"`
}

type Follow struct {
	FollowId   int64 `gorm:"not null"`
	FollowerId int64 `gorm:"not null"`
}
