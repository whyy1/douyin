package service

import (
	"douyin/dao"
)

func GetFavoriteList(userid int64) []dao.Video {
	videos := dao.FavoriteList(userid)
	return videos
}
