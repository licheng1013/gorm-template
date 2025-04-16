package tool

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"strings"
)

// 七牛云密钥 - 密钥查看: https://portal.qiniu.com/user/key
const accessKey = "your accessKey"
const secretKey = "your secretKey"

// 你的资源资源桶名称 - 桶查看: https://portal.qiniu.com/kodo/bucket
const bucket = "my-video-test"

// 获取token
func getToken() {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	MyLog.Println(upToken)
}

// DeleteFile 删除文件
func DeleteFile(key string) error {
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		UseHTTPS: false,
		Zone:     &storage.ZoneHuanan,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err := bucketManager.Delete(bucket, key)
	return err
}

// DeleteFileByUrl 根据url删除
func DeleteFileByUrl(url string) error {
	key := ParseKey(url)
	return DeleteFile(key)
}

// ParseKey 解析 url中的key
func ParseKey(url string) string {
	// 从url中解析出key
	index := strings.LastIndex(url, "/")
	key := url[index+1:]

	// 去除参数
	index = strings.Index(key, "?")
	if index != -1 {
		key = key[:index]
	}
	return key
}
