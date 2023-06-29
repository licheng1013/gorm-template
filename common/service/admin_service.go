package service

import (
	"common/app"
	"common/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var AdminService = adminService{}

// adminService 业务层
type adminService struct {
}

func (t adminService) db() *gorm.DB {
	return app.Db
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
func (t adminService) Delete(ids []int) {
	t.db().Delete(model.Admin{}, ids)
}
