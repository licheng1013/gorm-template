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
	Page int64 `json:"page" form:"page"`
	Size int64 `json:"size" form:"size"`
}
