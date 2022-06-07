package config

import (
	"encoding/json"
	"os"
)

//数据库配置参数
var DBuser string = "root"
var DBpassword string = "root"
var DBhost string = "127.0.0.1"
var DBport string = "3306"
var DBname string = "douyin"

var Conf aliyunOss

type aliyunOss struct {
	AccessKeyID     string `json:"AccessKeyID"`
	AccessKeySecret string `json:"AccessKeySecret"`
	Endpoint        string `json:"Endpoint"`
	Bucket          string `json:"Bucket"`
}

//阿里云OSS配置参数
func init() {
	file, _ := os.ReadFile("./config/config.json")
	_ = json.Unmarshal(file, &Conf)
}
