package util

import (
	"fmt"
	"os"

	"github.com/whyy1/douyin/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewBucket() *oss.Bucket {
	// 创建OSSClient实例。
	client, err := oss.New(config.Conf.Aliyun.Endpoint, config.Conf.Aliyun.AccessKeyID, config.Conf.Aliyun.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写存储空间名称
	bucket, err := client.Bucket(config.Conf.Aliyun.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return bucket
}
