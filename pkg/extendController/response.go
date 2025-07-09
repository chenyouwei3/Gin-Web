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

var ErrorCodeMap = map[int]ResponseMsg{
	//4000 自定义错误
	//参数错误
	4001: {ZhCn: "参数错误 - 空", EnUs: "Parameter Error - Empty"},
	4002: {ZhCn: "参数错误 - 绑定", EnUs: "Parameter Error - Binding"},
	4003: {ZhCn: "参数错误 - 不符合规则", EnUs: "Parameter error - does not comply with rules"},

	4010: {ZhCn: "身份信息不通过", EnUs: "Identity information not passed"},

	4040: {ZhCn: "资源不存在", EnUs: "The resource does not exist"},
	4090: {ZhCn: "数据重复", EnUs: "Data duplication"},

	4290: {ZhCn: "当前服务器过载,请稍后重试", EnUs: "The current server is overloaded, please try again later"},

	//5000自定义错误
	//数据库错误
	5100: {"添加失败", "Add failed"},
	5101: {ZhCn: "数据已存在", EnUs: "The data already exists"},

	5110: {"删除失败", "Delete failed"},

	5120: {"更新失败", "Modification failed"},

	5130: {"查询失败", "Query failed"},
}
