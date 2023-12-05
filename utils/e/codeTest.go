package e

var (
	Success         = NewError(0, "成功")
	ServerError     = NewError(10000000, "服务内部错误")
	ControllerError = NewError(510, "接收参数错误")
)
