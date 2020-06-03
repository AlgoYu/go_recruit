package common

// Response结构体
type Result struct {
	Code int
	Success bool
	Message string
	Data interface{}
}

// 构造返回类对象
func NewResult(code int, success bool, message string, data interface{}) *Result {
	return &Result{Code: code, Success: success, Message: message, Data: data}
}

// 成功
func Success(data interface{}) *Result {
	return &Result{Code: 200, Success: true, Message: "Success", Data: data}
}

// 失败
func Fail(message string) *Result {
	return &Result{Code: 500, Success: false, Message: message}
}