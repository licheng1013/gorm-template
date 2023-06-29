package app

import (
	"github.com/google/uuid"
	"strings"
)

// GetRandomStr 生成随机字符串
func GetRandomStr() string {
	uid, _ := uuid.NewUUID()
	return strings.ReplaceAll(uid.String(), "-", "")
}
