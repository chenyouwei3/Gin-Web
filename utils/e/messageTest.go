package e

import "fmt"

var codes = map[int]string{}

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
