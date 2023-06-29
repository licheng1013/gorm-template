package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	handlerFunc := cors.New(cors.Config{
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Authorization", "content-type"}, //此处设置非默认之外的请求头(自定义请求头),否则会出现跨域问题,必须制定名称
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowWildcard:    true,
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           24 * time.Hour,
	})

	return handlerFunc
}
