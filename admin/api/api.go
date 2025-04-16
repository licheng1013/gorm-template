package api

import (
	"common/middleware"
	"common/tool"
	"github.com/gin-gonic/gin"
)

func Init() {
	tool.MyLog.Println("注册中间件")
	tool.R.Use(gin.Logger())
	tool.R.Use(middleware.Cors())
	tool.R.Use(middleware.NewBaseAuth(middleware.AdminAuth{}))
	tool.R.Use(middleware.Recover)
	tool.MyLog.Println("注册接口")
	AdminApi{}.init(tool.R)
	FileApi{}.init(tool.R)
}
