package dao

import (
	"time"
)

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	VideoId    int64  `gorm:"not null"`
	UserId     int64  `gorm:"not null"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	User       User   `json:"user" gorm:"foreignkey:UserId"`
}

func GetComment(commentid int64) Comment {
	comment := Comment{
		Id: commentid,
	}
	db.First(&comment, "id = ?", commentid)
	return comment
}

func AddComment(text string, user User, userid int64, videoid int64) Comment {
	comment := Comment{
		User:       user,
		Content:    text,
		UserId:     userid,
		VideoId:    videoid,
		CreateDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	db.Create(&comment)
	return comment
}

func DeleteComment(commentid int64) {
	db.Where("id =?", commentid).Delete(&Comment{})
}

func CommentList(videoid int64) []Comment {
	list := []Comment{}

	db.Preload("User").Order("create_date desc").Where("video_id=?", videoid).Find(&list)
	return list
}
