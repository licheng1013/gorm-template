package api

import (
	"common/app"
	"common/model"
	"common/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	c := &AdminApi{}
	c.init(app.R) //这里需要引入你的gin框架的实例
}

func (t AdminApi) init(g *gin.Engine) {
	// 依次: 分页列表，单条，插入，修改，删除
	group := g.Group("/admin")
	group.GET("/list", t.list) //不设置限制条件的画默认查询所有
	group.GET("/one", t.one)
	group.POST("/insert", t.insert)
	group.POST("/update", t.update)
	group.POST("/delete", t.delete)
}

// AdminApi 控制器
type AdminApi struct {
	Page int   `form:"page,default=1"`
	Size int   `form:"size,default=10"`
	Ids  []int `form:"ids"`
}

// 分页列表
func (t AdminApi) list(c *gin.Context) {
	_ = c.Bind(&t)
	v := model.Admin{}
	_ = c.Bind(&v)
	c.JSON(http.StatusOK, model.OkData(service.AdminService.List(t.Page, t.Size, &v)))
}

// 根据主键Id查询记录
func (t AdminApi) one(c *gin.Context) {
	var v model.Admin
	_ = c.Bind(&v)
	c.JSON(http.StatusOK, model.OkData(service.AdminService.One(v.AdminId)))
}

// 修改记录
func (t AdminApi) update(c *gin.Context) {
	var v model.Admin
	_ = c.ShouldBindJSON(&v)
	service.AdminService.Update(v)
	c.JSON(http.StatusOK, model.OkMsg("修改成功！"))
}

// 插入记录
func (t AdminApi) insert(c *gin.Context) {
	var v model.Admin
	_ = c.ShouldBindJSON(&v)
	service.AdminService.Insert(v)
	c.JSON(http.StatusOK, model.OkMsg("插入成功！"))
}

// 根据主键删除
func (t AdminApi) delete(c *gin.Context) {
	_ = c.Bind(&t)
	service.AdminService.Delete(t.Ids)
	c.JSON(http.StatusOK, model.OkMsg("删除成功！"))
}
