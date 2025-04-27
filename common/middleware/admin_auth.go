package middleware

import "common/service"

// AdminAuth 管理员后台认证
type AdminAuth struct {
}

func (b AdminAuth) GetUserId(id any) int64 {
	user := service.AdminService.One(id)
	return user.Id
}
