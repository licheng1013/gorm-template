package tool

import (
	"fmt"
)

// NewServiceError 使用关键字 panic 抛出异常,否则无法识别到
func NewServiceError(text string) error {
	return &ServiceError{text}
}

type ServiceError struct {
	s string
}

func (e *ServiceError) Error() string {
	return e.s
}

// ErrorToString recover错误，转string
func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case *ServiceError:
		return v.Error()
	case ServiceError:
		return v.Error()
	default:
		MyLog.Println("系统错误:", v)
		return fmt.Sprint(v)
	}
}
