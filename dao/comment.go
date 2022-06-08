package dao

import (
	"fmt"
	"time"
)

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	VideoId    int64  //`gorm:"not null" gorm:"foreignkey:VideorId"`
	UserId     int64  `gorm:"not null"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	User       User   `json:"user" gorm:"foreignkey:UserId"`
}

func AddComment(text string, user User, userid int64, videoid int64) (Comment, error) {
	//对Comment表插入评论
	comment := Comment{
		User:       user,
		Content:    text,
		UserId:     userid,
		VideoId:    videoid,
		CreateDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Create(&comment).Error; err != nil {
		fmt.Println("评论插入失败", err)
		return comment, err
	}
	return comment, nil
}

func DeleteComment(commentid int64) error {
	//删除Comment表中评论
	if err := db.Debug().Where("id =?", commentid).Delete(&Comment{}).Error; err != nil {
		fmt.Println("评论删除失败", err)
		return err
	}
	return nil
}
func CommentList(videoid int64) ([]Comment, error) {
	list := []Comment{}
	if err := db.Debug().Preload("User").Order("create_date desc").Where("video_id=?", videoid).Find(&list).Error; err != nil {
		fmt.Println("评论列表获取失败", err)
		return list, err
	}
	return list, nil
}
