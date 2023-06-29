package api

import (
	"common/app"
	"common/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

// Init 注意包注册必须要在包内的第一个. 因为go是按字母顺序初始化包的
func Init() {
	log.Println("注册Admin-Api")
}

func init() {
	log.Println("注册中间件")
	// 跨域中间件必须在前面,否则会报错
	app.R.Use(gin.Logger())
	app.R.Use(middleware.Cors())
	app.R.Use(middleware.AuthAdmin)
	app.R.Use(middleware.Recover)
	// 中间件应该在路由注册之前注册, 包初始化顺序
	// 被引入的包优化执行文件中的, init 方法 -> 当前包 init -> 引入包方法的执行 -> 当前包方法的执行
}
