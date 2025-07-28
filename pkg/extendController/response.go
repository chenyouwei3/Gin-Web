package extendController

type Response struct {
	Code    int         `json:"code"` //自定义的http code
	Message ResponseMsg `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseMsg struct {
	ZhCn string `json:"zh-CN"`
	EnUs string `json:"en-US"`
}
