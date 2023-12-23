package api

import (
	"common/app"
	"common/model"
	"common/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	c := &FileApi{}
	c.init(app.R) //这里需要引入你的gin框架的实例
}

type FileApi struct {
}

func (t FileApi) init(g *gin.Engine) {
	group := g.Group("/file")
	group.POST("/upload", t.upload)
}

func (t FileApi) upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, app.Config.UploadPath+file.Filename)
	if err != nil {
		tool.AssertErr(err.Error())
	}
	c.JSON(http.StatusOK, model.OkMsg("上传成功"))
}
