package service

import (
	"douyin/dao"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
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

func SaveVideo(userid int64, playurl string, title string) (*dao.Video, error) {
	video := dao.Video{
		UserId:        userid,
		PlayUrl:       playurl,
		CoverUrl:      "http://y1-image.oss-cn-beijing.aliyuncs.com/image/2022/05/06/0593069c-f577-4bd2-9076-c07e3ad9904c.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Titile:        title,
	}
	video, err := dao.PublishVideo(&video)
	return &video, err
}

//文件上传到OSS
func PublishVideo(data *multipart.FileHeader, userid int64) error {
	file, _ := data.Open()
	bucket := dao.NewBucket()

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", userid, filename)
	path := fmt.Sprintf("%v/%v/%v/%v", time.Now().Year(), time.Now().Month(), time.Now().Day(), finalName)
	//savePath := filepath.Join(path, finalName)
	err := bucket.PutObject(path, file)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return err
}

func GetPubilshList(userid int64) []dao.Video {
	videos := dao.PublishList(userid)
	return videos
}
