package extendController

const (
	ErrQuery            = "query参数错误" //10001
	ErrBody             = "Body参数错误"  //10002
	ErrFormatConversion = "格式转化错误"    //10003
	ErrFormat           = "参数格式错误"
	ErrPaging           = "分页错误"
	ErrGetAll           = "查询所有数据错误"
)

// HTTP Code
const (
	Normal         = 200 // 无错误
	ParameterError = 400 // 参数错误
	NotFound       = 404 // 无资源错误
	ServerError    = 500 // 系统错误
	UnknownError   = 503 // 未知错误
	Unauthorized   = 401 // 未授权
)
