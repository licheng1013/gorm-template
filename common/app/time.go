package app

import (
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/datetime"
)

// GetTodayStartTime 今天开始时间 例如: 2020-01-01 00:00:00
func GetTodayStartTime() string {
	return datetime.GetNowDate() + " 00:00:00"
}

// GetTodayEndTime 今天结束时间 例如: 2020-01-01 23:59:59
func GetTodayEndTime() string {
	return datetime.GetNowDate() + " 23:59:59"
}

// ToInt 转换为int
func ToInt(str string) int64 {
	toInt, err := convertor.ToInt(str)
	if err != nil {
		return 0
	}
	return toInt
}
