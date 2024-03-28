package app

import (
	"testing"
)

func TestJwtToken(t *testing.T) {
	token, err := JwtUtil.CreateToken("1")
	if err != nil {
		tool.MyLog.Println("错误:" + err.Error())
	}
	tool.MyLog.Println(token)
	parseToken, err := JwtUtil.ParseToken(token)
	if err != nil {
		tool.MyLog.Println("错误:" + err.Error())
	}
	tool.MyLog.Println(parseToken)
}
