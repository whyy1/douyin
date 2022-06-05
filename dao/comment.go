package dao

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user" gorm:"foreignkey:UserId"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	VideoId    int64  //`gorm:"not null" gorm:"foreignkey:VideorId"`
}
