package app

import (
	"common/component"
	"github.com/gin-gonic/gin"
)

// R Gin实例
var R = gin.New()

// GlobalCache 全局缓存
var GlobalCache = component.NewCache()
