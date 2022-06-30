package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const (
	MaxFileSize     = 10 << 26 // 10 MB
	Endpoint        = "http://oss-cn-hangzhou.aliyuncs.com"
	AccessKeyId     = "L"
	AccessKeySecret = "v"
	BucketName      = "dousheng-bwcxgdz"
)

var (
	OssClient *oss.Client
)

func InitOSS() {
	var err error
	OssClient, err = oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		panic(err)
	}
}
