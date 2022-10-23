package dao

import (
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/whyy1/douyin/config"
)

var pool *redis.Pool

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Println("Redis配置读取失败: ", err)
	}
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive: 0, //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		//IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", config.REDIS_SOURCE)
		},
	}
}

func SetToken(id int64, token string) {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("Set", id, token)
	if err != nil {
		log.Println("redis:设置token错误，id", id, " faild :", err)
		log.Println("redis未连接")
		return
	}
}
func GetToken(id int64) (string, error) {
	c := pool.Get()
	defer c.Close()
	token, err := redis.String(c.Do("Get", id))
	if err != nil {
		log.Println("redis:获取token错误，id", id, " faild :", err)
		log.Println("redis未连接")
		return "", err
	}
	return token, nil
}
