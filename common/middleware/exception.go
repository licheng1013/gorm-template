package middleware

import (
	"bytes"
	"common/model"
	"common/tool"
	"io"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
)

// Recover 全局异常处理器
func Recover(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
	defer func() {
		if r := recover(); r != nil {
			tool.MyLog.Println("错误信息: ", r)
			// 对于不同的请求方式获取不同的传参,有助于定义参数问题
			if c.Request.Method == "POST" {
				// 如果是上传文件则跳过
				if strings.Contains(c.Request.Header.Get("Content-Type"),"multipart/form-data") {
					return
				}
				tool.MyLog.Println("请求Body:", string(body))
			} else if c.Request.Method == "GET" {
				tool.MyLog.Println("请求参数:", c.Request.URL.RawQuery)
			}
			debug.PrintStack() //打印错误堆栈信息
			c.JSON(http.StatusOK, model.Fail(tool.ErrorToString(r)))
			c.Abort() //终止后续接口调用，不加的话,还会继续接口后续代码
		}
	}()
	c.Next()
}
