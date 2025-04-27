package service

import (
	"common/model"
	"common/tool"
)

// 根据对象查询
func One[T any](e T) T {
	var one T
	err := tool.Db.Where(e).First(&one).Error
	tool.AssertError(err, "查询失败!")
	return one
}

func OneById[T any](id any) T {
	var one T
	err := tool.Db.First(&one, id).Error
	tool.AssertError(err, "查询失败!")
	return one
}

func Page[T any](e T, page, size int) model.PageData[T] {
	lists := make([]T, 0) // 结果
	w := tool.Db.Model(e).Where(e)
	var total int64 // 统计
	w.Count(&total)
	w.Order("create_time DESC").Offset((page - 1) * size).Limit(size).Find(&lists)
	return model.PageData[T]{
		List:  lists,
		Total: total,
	}
}

// e更新对象 w条件
func Update[T any](e, w T) {
	err := tool.Db.Model(w).Where(w).Updates(e).Error
	tool.AssertError(err, "更新失败!")
}

// 插入
func Insert[T any](e T) {
	err := tool.Db.Save(e).Error
	tool.AssertError(err, "插入失败!")
}

// 根据条件删除
func Delete[T any](w T) {
	err := tool.Db.Where(w).Delete(w).Error
	tool.AssertError(err, "删除失败!")
}

func DeleteById[T any](id any) {
	var m T
	err := tool.Db.Delete(m, id).Error
	tool.AssertError(err, "删除失败!")
}
