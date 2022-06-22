package service

import (
	"douyin/dao"
	"douyin/util"
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

func SaveVideo(userid int64, playurl string, coverurl string, title string) (*dao.Video, error) {
	video := dao.Video{
		UserId:        userid,
		PlayUrl:       playurl,
		CoverUrl:      coverurl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Titile:        title,
	}
	video, err := dao.PublishVideo(&video)
	return &video, err
}

//文件上传到OSS
func PublishVideo(data *multipart.FileHeader, userid int64) (string, string, error) {
	file, _ := data.Open()
	bucket := util.NewBucket()

	filename := time.Now().Format("15:04:05") + filepath.Ext(data.Filename)
	finalName := fmt.Sprintf("%d_%s", userid, filename)
	time := time.Now().Format("2006/01/02")
	path := fmt.Sprintf("%v/%v", time, finalName)

	err := bucket.PutObject(path, file)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	playurl := fmt.Sprintf("http://y1-douyin.oss-cn-hangzhou.aliyuncs.com/%v", path)

	//视频抽帧
	coverurl := playurl + "?x-oss-process=video/snapshot,t_500,f_jpg,w_0,h_0,m_fast"
	return playurl, coverurl, err
}

func GetPubilshList(userid int64) []dao.Video {
	videos := dao.PublishList(userid)
	return videos
}
