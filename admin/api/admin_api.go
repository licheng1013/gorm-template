package api

import (
	"common/model"
	"common/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t AdminApi) init(g *gin.Engine) {
	group := g.Group("/admin")
	group.GET("/list", t.list)
	group.GET("/one", t.one)
	group.POST("/insert", t.insert)
	group.POST("/update", t.update)
	group.POST("/delete", t.delete)
	group.POST("/login", t.login)
	group.GET("/getUserInfo", t.getUserInfo)
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
	c.JSON(http.StatusOK, model.OkData(service.AdminService.List(t.Page, t.Size, v)))
}

// 根据主键Id查询记录
func (t AdminApi) one(c *gin.Context) {
	var v model.Admin
	_ = c.Bind(&v)
	c.JSON(http.StatusOK, model.OkData(service.AdminService.One(v.Id)))
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
	var v model.Admin
	_ = c.ShouldBindJSON(&v)
	service.AdminService.Delete(v.Id)
	c.JSON(http.StatusOK, model.OkMsg("删除成功！"))
}

func (t AdminApi) login(c *gin.Context) {
	var v model.Admin
	_ = c.ShouldBindJSON(&v)
	// 登入成功,返回token
	c.JSON(http.StatusOK, model.OkData(service.AdminService.Login(v)))
}

func (t AdminApi) getUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, model.OkData(service.AdminService.GetUserInfo()))
}
