package util

import (
	"fmt"
	"log"
	"os"

	"github.com/whyy1/douyin/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewBucket() *oss.Bucket {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	// 创建OSSClient实例。
	client, err := oss.New(config.OSS_ENDPOINT, config.OSS_ACCESS_KEY_ID, config.OSS_ACCESS_KEY_SECRET)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写存储空间名称
	bucket, err := client.Bucket(config.OSS_BUCKET)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return bucket
}
