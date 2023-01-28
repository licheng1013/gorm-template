package dto

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// OkData 成功的数据返回
func OkData(data interface{}) Result {
	return Result{Code: 0, Data: data}
}

// OkMsg 成功的消息返回
func OkMsg(msg string) Result {
	return Result{Code: 0, Msg: msg}
}

// Fail 失败的消息返回
func Fail(msg string) Result {
	return Result{Code: -1, Msg: msg}
}
