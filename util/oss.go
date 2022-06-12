package util

import (
	"douyin/config"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewBucket() *oss.Bucket {
	// 创建OSSClient实例。
	client, err := oss.New(config.Conf.Endpoint, config.Conf.AccessKeyID, config.Conf.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写存储空间名称
	bucket, err := client.Bucket(config.Conf.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return bucket
}
