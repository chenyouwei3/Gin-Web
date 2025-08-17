package extendController

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
	RunLog *zap.Logger
}

func (b *BaseController) SendResponse(c *gin.Context, httpResponseCode, CustomCode int, msg ResponseMsg, data interface{}, err error) {
	if err != nil {
		msg.EnUs = fmt.Sprintf("%s : %v", msg.EnUs, err)
		b.RunLog.Error(msg.ZhCn + "/|^_^|/" + msg.EnUs) //定义日志输出格式
	} else {
		b.RunLog.Info(msg.ZhCn + "/|^_^|/" + msg.EnUs) //定义日志输出格式
	}

	c.JSON(httpResponseCode, Response{
		Code:    CustomCode, //自定义错误码
		Message: msg,        //错误信息还是正确信息
		Data:    data,       //数据
	})
}

// 成功
func (b *BaseController) SendSuccessResponse(c *gin.Context, data interface{}) {
	b.SendResponse(c, http.StatusOK, 2000, ResponseMsg{
		ZhCn: "请求成功",
		EnUs: "success",
	}, data, nil)
}

// 客户端自定义错误
func (b *BaseController) SendCustomResponseByFront(c *gin.Context, ZhCn, EnUs string, err error) {
	b.SendResponse(c, http.StatusBadRequest, 4000, ResponseMsg{
		ZhCn: ZhCn,
		EnUs: EnUs,
	}, nil, err)
}

// 服务端自定义错误
func (b *BaseController) SendCustomResponseByBacked(c *gin.Context, ZhCn, EnUs string, err error) {
	b.SendResponse(c, http.StatusOK, 5000, ResponseMsg{
		ZhCn: ZhCn,
		EnUs: EnUs,
	}, nil, err)
}

// 客户端参数错误
func (b *BaseController) SendParameterErrorResponse(c *gin.Context, CustomCode int, err error) {
	b.SendResponse(c, http.StatusBadRequest, CustomCode, ErrorCodeMap[CustomCode], nil, err)
}

// 服务端错误代码
func (b *BaseController) SendServerErrorResponse(c *gin.Context, CustomCode int, err error) {
	b.SendResponse(c, http.StatusOK, CustomCode, ErrorCodeMap[CustomCode], nil, err)
}

// 方法不允许405
func (b *BaseController) SendMethodNotAllowedResponse(c *gin.Context) {
	b.SendResponse(c, http.StatusMethodNotAllowed, 405, ResponseMsg{
		ZhCn: "方法不允许",
		EnUs: "Method not allow",
	}, nil, nil)
}

// 鉴权失败 401 Unauthorized
func (b *BaseController) SendUnauthorizedResponse(c *gin.Context, err error) {
	b.SendResponse(c, http.StatusUnauthorized, 4010, ErrorCodeMap[4010], nil, err)
}

// 请求过多429
func (b *BaseController) SendTooManyResponse(c *gin.Context, err error) {
	b.SendResponse(c, http.StatusTooManyRequests, 4290, ErrorCodeMap[4290], nil, err)
}
