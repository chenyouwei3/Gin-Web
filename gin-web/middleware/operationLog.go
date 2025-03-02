package middleware

import (
	"gin-web/initialize/mysql"
	"gin-web/initialize/runLog"
	"gin-web/models"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
	"time"
)

// 第一个参数是项目的前缀请求
// 第二个是需要过敏处理的参数
// 第三个是数据库连接

func OperationLog(target string, targets []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求链路
		linkStartTime := time.Now() // 记录请求开始时间
		c.Next()
		linkEndTime := time.Now()                           // 记录请求结束时间
		costTime := linkEndTime.Sub(linkStartTime).String() // 计算请求处理时间
		//响应区间
		Path, startTime := c.Request.URL.Path, time.Now()
		if !strings.HasPrefix(Path, target) || c.Request.Method == "GET" { //不存在的api不记录且GET方法不记录
			return
		}
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			runLog.ZapLog.Info("JSON deserialization failed")
		}
		username, _ := c.Get("user")
		operationLog := models.OperationLog{
			Username:  username.(string),
			Ip:        c.ClientIP(),
			Method:    c.Request.Method,
			Status:    c.Writer.Status(),
			Query:     c.Request.URL.RawQuery,
			Body:      string(body),
			Path:      Path,
			StartTime: startTime,
			UserAgent: c.Request.UserAgent(),
			CostTime:  costTime,
		}
		//脱敏处理
		for _, target := range targets {
			if Path == target {
				operationLog.Body, operationLog.Query = "***敏感信息已脱敏***", "***敏感信息已脱敏***"
			}
		}
		//可以根据需求自行修改
		go func() {
			err = mysql.DB.Create(&operationLog).Error
			if err != nil {
				runLog.ZapLog.Error(err.Error())
			}
		}()
	}
}
