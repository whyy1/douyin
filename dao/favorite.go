package dao

import "fmt"

type Favorite struct {
	UserId  int64 `gorm:"not null"`
	VideoId int64 `gorm:"not null"`
}

func AddFavorite(userid int64, videoid int64) (Favorite, error) {
	favorite := Favorite{userid, videoid}
	if err := db.Create(&favorite).Error; err != nil {
		fmt.Println("点赞失败", err)
	}
	return favorite, nil
}

func DeleteFavorite(userid int64, videoid int64) error {
	//删除Favorite表中记录
	if err := db.Debug().Where("user_id = ?", userid).Where("video_id = ?", videoid).Delete(&Favorite{}).Error; err != nil {
		fmt.Println("点赞失败", err)
	}
	return nil
}
func FavoriteList(userid int64) []Video {
	favoritelist := []Video{}
	arr := []int64{}
	db.Debug().Table("favorites").Select("video_id").Where("user_id = ?", userid).Scan(&arr)
	db.Debug().Preload("Author").Where("id IN ?", arr).Find(&favoritelist)
	return favoritelist
}
