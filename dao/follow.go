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

//传入user_id、touserid,删除follows表中记录
func DeleteFollow(userid int64, followid int64) {
	db.Where("follow_id = ?", followid).Where("follower_id = ?", userid).Delete(&Follow{})
}

//增加被关注用户的follow_count数以及关注用户的follower_count数
func AddFollowCount(userid int64, followid int64) {
	user := User{}
	touser := User{}

	db.Debug().First(&user, "id = ?", userid).Updates(User{FollowCount: user.FollowCount + 1})
	db.Debug().First(&touser, "id = ?", followid).Updates(User{FollowerCount: touser.FollowerCount + 1})
}

//减少被关注用户的follow_count数以及关注用户的follower_count数
func DeductFollowCount(userid int64, followid int64) {
	user := User{}
	touser := User{}

	db.Debug().First(&user, "id = ?", user.Id).Updates(User{FollowCount: user.FollowCount - 1})
	db.Debug().First(&touser, "id = ?", followid).Updates(User{FollowerCount: touser.FollowerCount - 1})
}

func FollowList(userid int64, authorid int64) []User {
	followerlist := []User{}
	follow := []int64{}
	arr := []int64{}

	db.Debug().Table("follows").Select("follow_id").Where("follower_id = ?", userid).Scan(&follow)
	db.Debug().Table("follows").Select("follow_id").Where("follower_id= ?", authorid).Scan(&arr)
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

	db.Debug().Table("follows").Select("follow_id").Where("follower_id = ?", userid).Scan(&follow)
	db.Debug().Table("follows").Select("follower_id").Where("follow_id= ?", authorid).Scan(&arr)
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
