package tool

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/datetime"
	"github.com/google/uuid"
	"log"
	"os"
	"strings"
)

var MyLog *log.Logger

func init() {
	MyLog = log.New(os.Stdout, "系统日志: ", log.Lshortfile+log.LstdFlags)
}

// 加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 获取Md5后的密码
func GetMd5Password(password, salt string) string {
	return Md5(password + salt)
}

// 今天开始时间 例如: 2020-01-01 00:00:00
func GetTodayStartTime() string {
	return datetime.GetNowDate() + " 00:00:00"
}

// 今天结束时间 例如: 2020-01-01 23:59:59
func GetTodayEndTime() string {
	return datetime.GetNowDate() + " 23:59:59"
}

// 转换为int
func ToInt(str string) int64 {
	toInt, err := convertor.ToInt(str)
	if err != nil {
		return 0
	}
	return toInt
}

// 生成随机字符串
func GetRandomStr() string {
	uid, _ := uuid.NewUUID()
	return strings.ReplaceAll(uid.String(), "-", "")
}
