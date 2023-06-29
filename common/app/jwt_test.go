package app

import (
	"log"
	"testing"
)

func TestJwtToken(t *testing.T) {
	token, err := JwtUtil.CreateToken("1")
	if err != nil {
		log.Println("错误:" + err.Error())
	}
	log.Println(token)
	parseToken, err := JwtUtil.ParseToken(token)
	if err != nil {
		log.Println("错误:" + err.Error())
	}
	log.Println(parseToken)
}
