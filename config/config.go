package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Conf Yaml

//数据库配置参数
type MysqlConf struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

//阿里云OSS配置参数
type AliyunConf struct {
	AccessKeyID     string `yaml:"AccessKeyID"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
	Endpoint        string `yaml:"Endpoint"`
	Bucket          string `yaml:"Bucket"`
}

type Yaml struct {
	Mysql  MysqlConf  `yaml:"mysql"`
	Aliyun AliyunConf `yaml:"aliyun"`
}

func init() {
	file, _ := ioutil.ReadFile("./config/config.yaml")
	err := yaml.Unmarshal(file, &Conf)
	if err != nil {
		log.Fatalf("yaml unmarshal: %v\n", err)
		return
	}
}
