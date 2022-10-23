package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/whyy1/douyin/controller"
)

func main() {
	r := setupRouter()
	setupLogs()

	if err := r.Run(); err != nil {
		log.Println("startup service failed, err:", err)
	}
	log.Println("服务启动成功！")
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

func setupLogs() {
	logFile, err := os.OpenFile("./logs/douyin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("日志打开成功！")
}
