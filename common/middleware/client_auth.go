package middleware

import "common/service"

// ClientAuth 客户端认证
type ClientAuth struct {
}

func (b ClientAuth) GetUserId(id interface{}) int64 {
	user := service.AdminService.One(id)
	return user.Id
}
