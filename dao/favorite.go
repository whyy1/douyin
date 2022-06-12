package dao

type Favorite struct {
	UserId  int64 `gorm:"not null"`
	VideoId int64 `gorm:"not null"`
}

func AddFavorite(userid int64, videoid int64) {
	favorite := Favorite{userid, videoid}

	db.Create(&favorite)
}

func DeleteFavorite(userid int64, videoid int64) {
	db.Where("user_id = ?", userid).Where("video_id = ?", videoid).Delete(&Favorite{})
}

func FavoriteList(userid int64) []Video {
	favoritelist := []Video{}
	arr := []int64{}

	db.Table("favorites").Select("video_id").Where("user_id = ?", userid).Scan(&arr)
	db.Preload("Author").Where("id IN ?", arr).Find(&favoritelist)
	return favoritelist
}
