package model

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func OkData(data any) Result {
	return Result{Code: 0, Data: data}
}
func OkMsg(msg string) Result {
	return Result{Code: 0, Msg: msg}
}

func Fail(msg string) Result {
	return Result{Code: -1, Msg: msg}
}

type PageVo struct {
	Page int `json:"page" form:"page,default=1"`
	Size int `json:"size" form:"size,default=10"`
}

type PageData[T any] struct {
	Total int64 `json:"total"`
	List  []T   `json:"list"`
}
