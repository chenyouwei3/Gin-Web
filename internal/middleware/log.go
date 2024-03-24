package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"loopy-manager/initialize/global"
	"loopy-manager/internal/model"
	"time"
)

func OperationLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		path := c.Request.URL.Path      //请求路径
		query := c.Request.URL.RawQuery //query参数
		endTime := time.Now()
		costTime := endTime.Sub(startTime).Milliseconds()
		//user, _ := c.Get("user")
		operationLog := model.OperationLog{
			Id:        global.LogSnowFlake.Generate().Int64(),
			Username:  " ",
			Ip:        c.ClientIP(),
			Method:    c.Request.Method,
			Query:     query,
			Path:      path,
			Status:    c.Writer.Status(),
			StartTime: startTime,
			TimeCost:  costTime,
			UserAgent: c.Request.UserAgent(),
			Errors:    c.Errors.ByType(gin.ErrorTypePrivate).String(),
		}

		res := global.LogTable.Debug().Create(operationLog)
		if res.Error != nil {
			logrus.Error("中间件日志记录失败:", res.Error)
		}
		c.Next()
	}
}
