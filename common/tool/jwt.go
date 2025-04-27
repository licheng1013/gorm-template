package tool

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JwtUtil = NewJwt("你的密钥")

type Jwt struct {
	SecretKey []byte
}

func NewJwt(secretKey string) *Jwt {
	return &Jwt{SecretKey: []byte(secretKey)}
}

// CreateTokenWithTimeOut 由传过来的时间自定义
func (j Jwt) CreateTokenWithTimeOut(ID interface{}, date time.Duration) (string, error) {
	expirationTime := time.Now().Add(date)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		//Issuer:    "WebToken", //可空
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ID:       fmt.Sprint(ID),
		//Subject:   "MyInfo",  //可空
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// CreateToken 生成一个3天的Token
func (j Jwt) CreateToken(ID interface{}) (string, error) {
	return j.CreateTokenWithTimeOut(ID, time.Hour*24*3)
}

// CreateTokenEasy  生成一个3天的Token
func (j Jwt) CreateTokenEasy(ID interface{}) string {
	token, err := j.CreateToken(ID)
	AssertError(err, "生成Token失败")
	return token
}

// ParseToken 解析并验证JWT，并检查是否过期
func (j Jwt) ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("验证失败")
		}
		return j.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 如果token有效，将其断言为自定义的声明结构体，并返回其中的信息
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
