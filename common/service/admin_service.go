package service

import (
	"common/model"
	"common/tool"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var AdminService = adminService{}

// adminService 业务层
type adminService struct {
}

func (t adminService) db() *gorm.DB {
	return tool.Db
}

func (t adminService) List(page, size int, v model.Admin) any {
	return Page(v, page, size)
}

func (t adminService) Update(v model.Admin) {
	Update(v, v)
}

func (t adminService) Insert(v model.Admin) {
	Insert(v)
}

func (t adminService) Delete(id int64) {
	tool.AssertErr("测试不允许删除")
	DeleteById[model.Admin](id)
}

func (t adminService) Login(v model.Admin) interface{} {
	// 查询用户是否存在
	var adminDb model.Admin
	t.db().Where(model.Admin{UserName: v.UserName}).First(&adminDb)
	if adminDb.Id == 0 || adminDb.Password != tool.GetMd5Password(v.Password, adminDb.Salt) {
		tool.AssertErr("密码或账号错误")
	}
	return gin.H{"token": tool.JwtUtil.CreateTokenEasy(adminDb.Id)}
}

func (t adminService) GetAdminId() int64 {
	return tool.LocalData.GetCtx().(int64)
}

func (t adminService) GetUserInfo() interface{} {
	adminId := t.GetAdminId()
	return AdminDto{
		UserId:   adminId,
		Username: "管理员",
		RealName: "管理员",
	}
}

type AdminDto struct {
	// userId
	UserId int64 `json:"userId"`
	// 用户名
	Username string `json:"username"`
	// 真实名字
	RealName string `json:"realName"`
	// 头像
	Avatar string `json:"avatar"`
	// 介绍
	Desc string `json:"desc"`
}
