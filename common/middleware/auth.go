package middleware

import (
	"common/app"
	"common/model"
	"common/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const Authorization = "Authorization"

type BaseAuth interface {
	GetUserId(id interface{}) int64
}

func NewBaseAuth(auth BaseAuth) func(c *gin.Context) {
	impl := BaseAuthImpl{}
	impl.BaseAuth = auth
	return impl.Auth
}

// BaseAuthImpl 采用组合的方式实现接口
type BaseAuthImpl struct {
	BaseAuth
}

func isExclude(path string, excludePath []string) bool {
	// 这里的 path 是没有去出?后面的参数的,所以需要处理一下
	if strings.Contains(path, "?") {
		path = strings.Split(path, "?")[0]
	}
	//排除路径处理
	for _, item := range excludePath {
		var ePath = item
		if path == ePath {
			return true
		}
		// 判断尾部是否有/*
		if strings.HasSuffix(ePath, "/*") {
			k := strings.TrimSuffix(ePath, "/*")
			if strings.HasPrefix(path, k) {
				tool.MyLog.Println("排除路径:", k, "请求路径:", path)
				return true
			}
		}
	}
	return false
}

// Auth 管理员后台处理认证器
func (b BaseAuthImpl) Auth(c *gin.Context) {
	path := c.Request.RequestURI
	if isExclude(path, app.Config.ExcludePath) {
		return
	}
	//登入认证
	token := c.Request.Header.Get(Authorization)
	if token == "" {
		tool.MyLog.Println("没有登入!")
		c.JSON(http.StatusOK, model.Fail("没有登入!"))
		c.Abort() //终止访问
		return
	}
	// 验证用户是否存在
	data, err := app.JwtUtil.ParseToken(token)
	if err != nil {
		// -10 表示登入失效了,需要重新登入
		tool.MyLog.Println("登入失效!")
		c.JSON(http.StatusOK, model.Result{Code: -10, Msg: "登入失效!"})
		c.Abort() //终止访问
		return
	}
	userId := b.GetUserId(data.ID)
	if userId == 0 {
		tool.AssertErr("用户不存在")
	}
	// 保存用户信息到上下文中
	app.LocalData.SaveCtx(userId)
	c.Next()
	app.LocalData.ClearCtx()
}
