package controller

import (
	"gin-web/internal/manage/models"
	"gin-web/internal/manage/types"
	"gin-web/pkg"
	"gin-web/pkg/extendController"

	"github.com/gin-gonic/gin"
)

type LogHandlerController struct {
	extendController.BaseController
}

func (l *LogHandlerController) GetListByOperation() gin.HandlerFunc {
	return func(c *gin.Context) {
		//接收参数
		logReq := &types.LogByOperationGetListReq{
			Account:   c.Query("account"),
			CurrPage:  c.DefaultQuery("currPage", "1"),
			PageSize:  c.DefaultQuery("pageSize", "10"),
			StartTime: c.Query("startTime"),
			EndTime:   c.Query("endTime"),
		}
		skip, limit, err := pkg.GetPage(logReq.CurrPage, logReq.PageSize)
		if err != nil {
			l.SendServerErrorResponse(c, 5131, err)
			return
		}
		//DB操作
		logDB := models.OperationLog{
			Account: logReq.Account,
		}
		resDB, total, err := logDB.GetList(skip, limit, logReq.StartTime, logReq.EndTime)
		if err != nil {
			l.SendServerErrorResponse(c, 5130, err)
			return
		}
		l.SendSuccessResponse(c, types.LogByOperationGetListResp{
			Logs:  resDB,
			Total: total,
		})
	}
}

func (l *LogHandlerController) GetListByRun() gin.HandlerFunc {
	return func(c *gin.Context) {
		//接收参数
		//logReq := &types.LogByOperationGetListReq{}

	}
}
