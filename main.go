package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whyy1/douyin/controller"
)

func main() {
	r := setupRouter()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}


func setupRouter() *gin.Engine {
	r := gin.Default()
	apiRouter := r.Group("/douyin")

	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.GET("/publish/list/", controller.PublishList)

	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)

	return r
}
