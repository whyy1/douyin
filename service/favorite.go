package service

import (
	"github.com/whyy1/douyin/dao"
)

func FavoriteAction(actiontype string, userid int64, videoid int64) {
	if actiontype == "1" {
		//增加Favorite表中记录
		dao.AddFavorite(userid, videoid)
		//根据video_id,增加对应的Video的点赞数
		dao.AddFavoriteCount(videoid)
	} else if actiontype == "2" {
		//删除Favorite表中记录
		dao.DeleteFavorite(userid, videoid)
		//根据video_id,减少对应的Video的点赞数
		dao.DeductFavoriteCount(videoid)
	}
}
func GetFavoriteList(userid int64) []dao.Video {
	videos := dao.FavoriteList(userid)
	return videos
}
