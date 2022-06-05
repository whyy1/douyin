package service

import (
	"douyin/dao"
)

type VideoListResponse struct {
	Response
	VideoList []dao.Video `json:"video_list"`
}
type Video struct {
	VideoId       int64  `json:"id,omitempty"`
	Author        User   `json:"author" gorm:"foreignKey:UserId;references:UserId;"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count" gorm:"force:force"`
	CommentCount  int64  `json:"comment_count" gorm:"force:force"`
	IsFavorite    bool   `json:"is_favorite"`
	UserId        int64  `gorm:"not null"`
	Titile        string `json:"title,omitempty"`
	CreateDate    int64  `gorm:"autoCreateTime"`
}

func PublishVideo(param *Video) *dao.Video {

}

func GetPubilshList(userid int64) []dao.Video {
	videos := dao.PublishList(userid)
	return videos
}
