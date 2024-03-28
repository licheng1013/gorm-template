package tool

import (
	"github.com/duke-git/lancet/v2/strutil"
	"unicode/utf8"
)

// AssertErrList 断言多个错误列表...
func AssertErrList(err ...error) error {
	for _, e := range err {
		if e != nil {
			return e
		}
	}
	return nil
}

// AssertErr 断言字符串
func AssertErr(errInfo string) {
	panic(NewServiceError(errInfo))
}

// AssertErrWithErrInfo  断言字符串
func AssertErrWithErrInfo(err error, errInfo string) {
	if err != nil {
		MyLog.Println("错误: ", err)
		panic(NewServiceError(errInfo))
	}
}

// AssertStr   断言字符串
func AssertStr(str string, errInfo string) {
	if strutil.IsBlank(str) {
		panic(NewServiceError(errInfo))
	}
}

// AssertNumber 断言数字浮点数是否为空
func AssertNumber[T int | int64 | int32 | uint | uint16 | uint32](num T, errInfo string) {
	if num == 0 {
		panic(NewServiceError(errInfo))
	}
}

// AssertTrue 为真则抛出异常
func AssertTrue(b bool, errInfo string) {
	if b {
		panic(NewServiceError(errInfo))
	}
}

// AssertLen 断言长度是否等于某长度
func AssertLen(arr string, length int, errInfo string) {
	if utf8.RuneCountInString(arr) != length {
		panic(NewServiceError(errInfo))
	}
}

// AssertNotLen 断言不在范围内的长度, 如果小于min或者大于max则抛出异常
func AssertNotLen(arr string, min, max int, errInfo string) {
	if strutil.IsBlank(arr) {
		panic(NewServiceError(errInfo))
	}
	if utf8.RuneCountInString(arr) < min || utf8.RuneCountInString(arr) > max {
		panic(NewServiceError(errInfo))
	}
}

// AssertGtLenOrEmpty  大于长度或为空则抛出异常, 如: length = 12 , 那么 13 及以上异常
func AssertGtLenOrEmpty(arr string, length int, errInfo string) {
	if strutil.IsBlank(arr) {
		panic(NewServiceError(errInfo))
	}
	if utf8.RuneCountInString(arr) == 0 || utf8.RuneCountInString(arr) > length {
		panic(NewServiceError(errInfo))
	}
}
