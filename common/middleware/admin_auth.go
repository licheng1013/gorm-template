package middleware

import (
	"common/app"
	"common/model"
	"common/service"
	"common/tool"
	"github.com/gin-gonic/gin"

	"net/http"
	"strings"
)

func init() {
	//log.Println("注册认证中间件!")
	//app.R.Use(Auth)
}

const AuthorizationAdmin = "Authorization"

// AuthAdmin 管理员后台处理认证器
func AuthAdmin(c *gin.Context) {
	path := c.Request.RequestURI
	//排除路径处理
	for _, item := range app.Config.ExcludePath {
		// 排除路径例如： /test/** , 排除匹配 /test/ 的所有拦截
		var ePath = item
		//检查是否存在**号
		if strings.HasSuffix(ePath, "**") || path == ePath {
			var prefix = strings.ReplaceAll(ePath, "*", "")
			if strings.HasPrefix(path, prefix) {
				return
			}
		}
	}
	//登入认证
	token := c.Request.Header.Get(AuthorizationAdmin)
	if token == "" {
		c.JSON(http.StatusOK, model.Fail("没有登入!"))
		c.Abort() //终止访问
		return
	}

	// 验证用户是否存在
	data, err := app.JwtUtil.ParseToken(token)
	if err != nil {
		// -10 表示登入失效了,需要重新登入
		c.JSON(http.StatusOK, model.Result{Code: -10})
		c.Abort() //终止访问
		return
	}
	user := service.AdminService.One(data.ID)
	// TODO 此处应该是查询管理员账号
	if user.Id == 0 {
		tool.AssertErr("用户不存在")
	}
	// 保存用户信息到上下文中
	app.LocalData.SaveCtx(user.Id)
	c.Next()
	app.LocalData.ClearCtx()
}
