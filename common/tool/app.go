package tool

import (
	"github.com/gin-gonic/gin"
)

// R Gin实例
var R = gin.New()

// GlobalCache 全局缓存
var GlobalCache = NewCache()
