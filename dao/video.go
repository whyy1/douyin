package dao

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
