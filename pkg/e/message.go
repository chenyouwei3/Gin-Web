package e

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

var MsgFlags = map[int]string{
	SUCCESS:              "ok",
	ERROR:                "fail",
	INVALID_PARAMS:       "请求参数错误",
	ControllerError:      "控制层错误",
	ServerError:          "逻辑层错误",
	ParameterError:       "接收参数错误",
	ParameterStructError: "绑定结构体参数错误",
	CreateError:          "创建数据失败",
}
