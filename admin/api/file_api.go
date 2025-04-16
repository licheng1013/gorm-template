package api

import (
	"common/model"
	"common/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FileApi struct {
}

func (t FileApi) init(g *gin.Engine) {
	group := g.Group("/file")
	group.POST("/upload", t.upload)
}

func (t FileApi) upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, tool.Config.UploadPath+file.Filename)
	if err != nil {
		tool.AssertErr(err.Error())
	}
	c.JSON(http.StatusOK, model.OkData(tool.Config.UploadUrl+file.Filename))
}
