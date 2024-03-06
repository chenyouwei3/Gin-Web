package utils

import (
	"net/http"
)

// Response 响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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
