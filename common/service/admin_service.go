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

// List 分页列表
func (t adminService) List(page, size int, v *model.Admin) gin.H {
	lists := make([]model.Admin, 0) // 结果
	t.db().Model(&model.Admin{}).Where(&v).Order("").Offset((page - 1) * size).Limit(size).Find(&lists)
	var total int64 // 统计
	t.db().Model(&v).Where(&v).Count(&total)
	h := gin.H{"list": lists, "total": total}
	return h
}

// One 根据主键查询
func (t adminService) One(id interface{}) (v model.Admin) {
	t.db().Find(&v, id)
	return
}

// Update 修改记录
func (t adminService) Update(v model.Admin) {
	t.db().Model(&v).Updates(v)
}

// Insert 插入记录
func (t adminService) Insert(v model.Admin) {
	t.db().Save(&v)
}

// Delete 根据主键删除
func (t adminService) Delete(id int64) {
	tool.AssertErr("测试不允许删除")
	t.db().Delete(model.Admin{}, id)
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
