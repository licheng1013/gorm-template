package tool

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestLen(t *testing.T) {
	AssertNotLen("HelloHello", 0, 10, "长度过长")
	//AssertNotLen("Hello World NiHao", 0, 10, "长度过长")
	//AssertNotLen("你好你好你好你好", 0, 10, "长度过长")
	fmt.Println(utf8.RuneCountInString("你好aa"))
}
