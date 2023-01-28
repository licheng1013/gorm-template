package api

import (
	"gin-web-template/common"
	"gin-web-template/middleware"
	"gin-web-template/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 通过包初始化注入
func init() {
	api := IndexApi{}
	group := common.R.Group("/index")
	group.GET("", api.index)
	group.GET("err", api.err)
}

type IndexApi struct {
}

// 分页列表
func (t IndexApi) index(c *gin.Context) {
	log.Println("收到请求!")
	c.JSON(http.StatusOK, model.OkData("HelloWorld"))
}

// 分页列表
func (t IndexApi) err(c *gin.Context) {
	panic(middleware.NewServiceError("测试异常"))
}
