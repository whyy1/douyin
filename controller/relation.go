package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	followid, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actiontype := c.Query("action_type")
	user := User{}
	touser := User{}

	if err := db.First(&user, "token =?", token).Error; err == nil {
		if actiontype == "1" {
			follow := Follow{FollowId: followid, FollowerId: user.Id}
			//根据user_id、touserid,增加follows表中记录
			if err := db.Create(&follow).Error; err != nil {
				fmt.Println(err)
			}
			//增加被关注用户的follow_count数以及关注用户的follower_count数
			db.Debug().First(&user, "id = ?", user.Id).Updates(User{FollowCount: user.FollowCount + 1})
			db.Debug().First(&touser, "id = ?", followid).Updates(User{FollowerCount: touser.FollowerCount + 1})
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "关注成功"})
		} else if actiontype == "2" {
			//根据user_id、touserid,减少follows表中记录
			if err := db.Where("follow_id = ?", followid).Where("follower_id = ?", user.Id).Delete(&Follow{}).Error; err != nil {
				fmt.Println("取消点赞失败", err)
			}
			//减少被关注用户的follow_count数以及关注用户的follower_count数
			db.Debug().First(&user, "id = ?", user.Id).Updates(User{FollowCount: user.FollowCount - 1})
			db.Debug().First(&touser, "id = ?", followid).Updates(User{FollowerCount: touser.FollowerCount - 1})
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "取消关注成功"})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// 关注列表
func FollowList(c *gin.Context) {
	//通过token拿到当前登录用户id
	//查看作者id为user_id
	token := c.Query("token")
	authorid := c.Query("user_id")
	followerlist := []User{}
	follow := []int64{}
	arr := []int64{}
	user := User{}

	db.First(&user, "token =?", token)
	db.Debug().Table("follows").Select("follow_id").Where("follower_id = ?", user.Id).Scan(&follow)
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
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: followerlist,
	})
}

// 粉丝列表
func FollowerList(c *gin.Context) {
	//通过token拿到当前登录用户id
	//查看作者id为user_id
	token := c.Query("token")
	authorid := c.Query("user_id")
	followerlist := []User{}
	follow := []int64{}
	arr := []int64{}
	user := User{}

	db.First(&user, "token =?", token)
	db.Debug().Table("follows").Select("follow_id").Where("follower_id = ?", user.Id).Scan(&follow)
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
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: followerlist,
	})
}
