package utils

import (
	"net/http"
)

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
