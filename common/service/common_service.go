package service

import (
	"common/model"
	"common/tool"
)

// 根据对象查询
func One[T any](e T) {
	var one T
	err := tool.Db.Where(e).First(&one).Error
	tool.AssertErrWithErrInfo(err, "查询失败!")
}

func Page[T any](e T, page, size int) model.PageResult[T] {
	lists := make([]T, 0) // 结果
	w := tool.Db.Model(e).Where(e)
	var total int64 // 统计
	w.Count(&total)
	w.Order("create_time DESC").Offset((page - 1) * size).Limit(size).Find(&lists)
	return model.PageResult[T]{
		List:  lists,
		Total: total,
	}
}

// e更新对象 w条件
func Update[T any](e, w T) {
	err := tool.Db.Model(w).Where(w).Updates(e).Error
	tool.AssertErrWithErrInfo(err, "更新失败!")
}

// 插入
func Insert[T any](e T) {
	err := tool.Db.Save(e).Error
	tool.AssertErrWithErrInfo(err, "插入失败!")
}

// 根据条件删除
func Delete[T any](w T) {
	err := tool.Db.Where(w).Delete(w).Error
	tool.AssertErrWithErrInfo(err, "删除失败!")
}

func DeleteById[T any](id any) {
	var m T
	err := tool.Db.Delete(m, id).Error
	tool.AssertErrWithErrInfo(err, "删除失败!")
}
