package dao

type Follow struct {
	FollowId   int64 `gorm:"not null"`
	FollowerId int64 `gorm:"not null"`
}

//传入user_id、touserid,增加follows表中记录
func AddFollow(userid int64, followid int64) {
	follow := Follow{followid, userid}

	db.Create(&follow)
}

func DeleteFollow(userid int64, followid int64) {
	db.Where("follow_id = ?", followid).Where("follower_id = ?", userid).Delete(&Follow{})
}

func FollowList(userid int64, authorid int64) []User {
	followerlist := []User{}
	follow := []int64{}
	arr := []int64{}

	db.Table("follows").Select("follow_id").Where("follower_id = ?", userid).Scan(&follow)
	db.Table("follows").Select("follow_id").Where("follower_id= ?", authorid).Scan(&arr)
	db.Where("id IN ?", arr).Find(&followerlist)

	for i := range followerlist {
		for _, j := range follow {
			if followerlist[i].Id == j {
				followerlist[i].IsFollow = true
				break
			}
		}
	}
	return followerlist
}

func FollowerList(userid int64, authorid int64) []User {
	followerlist := []User{}
	follow := []int64{}
	arr := []int64{}

	db.Table("follows").Select("follow_id").Where("follower_id = ?", userid).Scan(&follow)
	db.Table("follows").Select("follower_id").Where("follow_id= ?", authorid).Scan(&arr)
	db.Where("id IN ?", arr).Find(&followerlist)

	for i := range followerlist {
		for _, j := range follow {
			if followerlist[i].Id == j {
				followerlist[i].IsFollow = true
				break
			}
		}
	}
	return followerlist
}
