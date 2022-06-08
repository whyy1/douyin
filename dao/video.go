package dao

import (
	"douyin/config"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

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

func newvideo(param *Video) Video {
	video := Video{
		Id:            param.Id,
		Author:        param.Author,
		PlayUrl:       param.PlayUrl,
		CoverUrl:      param.CoverUrl,
		FavoriteCount: param.FavoriteCount,
		CommentCount:  param.CommentCount,
		Favorite:      param.Favorite,
		IsFavorite:    param.IsFavorite,
		UserId:        param.UserId,
		Titile:        param.Titile,
		CreateDate:    param.CreateDate,
	}
	return video
}
func PublishVideo(param *Video) (Video, error) {
	video := newvideo(param)
	err := db.Create(&video).Error
	return video, err
}
func GetVideoList() []Video {
	videos := []Video{}
	db.Debug().Preload("Author").Order("create_date desc").Find(&videos)

	return videos
}

func VideoFavorite(userid int64) []int64 {
	favorites := []int64{}
	db.Debug().Table("favorites").Select("video_id").Where("user_id = ?", userid).Scan(&favorites)
	return favorites
}

func Videofollow(userid int64) []int64 {
	follow := []int64{}
	db.Debug().Table("follows").Select("follow_id").Where("follower_id = ?", userid).Scan(&follow)
	return follow
}

func PublishList(userid int64) []Video {
	publishlist := []Video{}
	db.Debug().Preload("Author").Where("user_id =?", userid).Find(&publishlist)
	return publishlist
}

func NewBucket() *oss.Bucket {
	// 创建OSSClient实例。
	client, err := oss.New(config.Conf.Endpoint, config.Conf.AccessKeyID, config.Conf.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写存储空间名称
	bucket, err := client.Bucket(config.Conf.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return bucket
}

//增加视频评论数
func AddCommentCount(videoid int64) error {
	video := Video{}

	if err := db.Debug().First(&video, "id=?", videoid).Update("comment_count", video.CommentCount+1).Error; err != nil {
		fmt.Println("评论数增加失败", err)
		return err
	}

	return nil
}

//减少视频评论数
func DeductCommentCount(videoid int64) error {
	video := Video{}

	if err := db.Debug().First(&video, "id=?", videoid).Update("comment_count", video.CommentCount-1).Error; err != nil {
		fmt.Println("评论数减少失败", err)
		return err
	}
	return nil
}

//传入video_id,增加Video的点赞数
func AddFavoriteCount(videoid int64) {
	video := Video{}

	db.Debug().First(&video, "id = ?", videoid).Updates(Video{FavoriteCount: video.FavoriteCount + 1})
}

//传入video_id,减少Video的点赞数
func DeductFavoriteCount(videoid int64) {
	video := Video{}
	db.Debug().First(&video, "id=?", videoid).Updates(Video{FavoriteCount: video.FavoriteCount - 1})
}
