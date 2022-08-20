package dao

import (
	"fmt"
	"log"

	"github.com/whyy1/douyin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	//fmt.Println(config.DB_SOURCE)
	db, err = gorm.Open(mysql.Open(config.DB_SOURCE), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	sqlDB.SetMaxIdleConns(10)     //连接池最大允许的空闲连接数
	sqlDB.SetMaxOpenConns(100)    //设置数据库连接池最大连接数
	sqlDB.SetConnMaxLifetime(200) //设置数据库连接池可重用链接得最大时间长度
	fmt.Println("Mysql链接成功！")
}
