package app

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// GetMd5Password 获取Md5后的密码
func GetMd5Password(password, salt string) string {
	return Md5(password + salt)
}
