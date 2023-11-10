package utils

import (
	"fmt"
	"net/http"
)

// 自定义错误状态码
var codes = map[int]string{}

// Response 响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Error 错误响应体
type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func SuccessMess(message string, data interface{}) Response {
	return Response{
		http.StatusOK,
		message,
		data,
	}
}

func ErrorMess(message string, data interface{}) Response {
	return Response{
		http.StatusInternalServerError,
		message,
		data,
	}
}
