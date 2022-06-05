package dao

type Favorite struct {
	UserId  int64  `gorm:"not null"`
	VideoId string `gorm:"not null"`
}
