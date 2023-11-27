package middleware

import "testing"

// 编写排除路径测试代码
func TestIsExclude(t *testing.T) {
	var excludePath = []string{
		"/admin/*",
		"/user/admin",
	}
	if !isExclude("/admin/login", excludePath) {
		t.Error("排除路径测试失败")
	}
	if !isExclude("/user/admin", excludePath) {
		t.Error("排除路径测试失败")
	}
	var all = []string{
		"/*",
	}
	if !isExclude("/admin/login", all) {
		t.Error("排除路径测试失败")
	}
}
