package dao

type Favorite struct {
	UserId  int64  `gorm:"not null"`
	VideoId string `gorm:"not null"`
}

func FavoriteList(userid int64) []Video {
	favoritelist := []Video{}
	arr := []int64{}
	db.Debug().Table("favorites").Select("video_id").Where("user_id = ?", userid).Scan(&arr)
	db.Debug().Preload("Author").Where("video_id IN ?", arr).Find(&favoritelist)
	return favoritelist
}
