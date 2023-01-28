package common

import (
	"gin-web-template/middleware"
	"github.com/gin-gonic/gin"
)

var R = gin.New()

// 全局中间件的初始化
func init() {
	R.Use(middleware.Cors())
	R.Use(middleware.Recover)
}
