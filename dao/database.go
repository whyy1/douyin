package dao

import (
	"douyin/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBuser, config.DBpassword, config.DBhost, config.DBport, config.DBname)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
}
