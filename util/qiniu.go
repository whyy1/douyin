package util

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/whyy1/douyin/config"
)

type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

func NewUpToken() string {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	putPolicy := storage.PutPolicy{
		Scope: config.QN_BUCKET,
	}
	mac := qbox.NewMac(config.QN_ACCESS_KEY, config.QN_SECRET_KEY)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

func PutFile(upToken string, path string, data *multipart.FileHeader) error {

	cfg := storage.Config{}
	// 空间对应的机房
	//cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	//cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	//cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	file, _ := data.Open()

	ret := MyPutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, data.Size, &putExtra)
	if err != nil {
		fmt.Println(err)
		fmt.Println("上传失败")
		return err
	}
	fmt.Println(ret.Bucket, ret.Key, ret.Fsize, ret.Hash, ret.Name)
	return nil
}
