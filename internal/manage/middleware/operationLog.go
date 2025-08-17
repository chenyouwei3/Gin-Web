package middleware

import (
	"bytes"
	"encoding/json"
	"gin-web/internal/manage/models"
	"io"
	"time"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 日志协程需要的通道
var operationLogChan = make(chan models.OperationLog, 10)

var (
	skipMap = map[string]struct{}{ //跳过日志记录(这里的不需要进行记录)
	}
	targetMap = map[string]struct{}{ //数据脱敏
		"/login": {}}
)

func OperationLog(zapLog *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime, Path := time.Now(), c.Request.URL.Path //请求链路---记录请求开始时间
		//不存在的api不记录且GET方法不记录
		if _, skipOK := skipMap[Path]; skipOK || c.Request.Method == "GET" {
			return
		}
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			zapLog.Warn("读取请求体失败:", zap.Error(err))
		} else {
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 重置
		}

		c.Next()
		costTime := time.Now().Sub(startTime).String() // 计算请求处理时间
		operationLog := models.OperationLog{
			Ip:        c.ClientIP(),
			Method:    c.Request.Method,
			Status:    c.Writer.Status(),
			Query:     c.Request.URL.RawQuery,
			Body:      string(bodyBytes),
			Path:      Path,
			StartTime: time.Now(),
			UserAgent: c.Request.UserAgent(),
			CostTime:  costTime,
		}
		//login情况特殊处理
		if Path == "/login" {
			var login struct {
				Account  string `json:"account"`
				Password string `json:"password"`
			}
			err = json.Unmarshal(bodyBytes, &login)
			if err != nil {
				zapLog.Error("日志中间件解码失败")
			}
			operationLog.Account = login.Account
		} else {
			operationLog.Account = "login.Account"
		}
		if _, ok := targetMap[Path]; ok {
			operationLog.Body, operationLog.Query = "***敏感信息已脱敏***", "***敏感信息已脱敏***"
		}
		select {
		case operationLogChan <- operationLog:
		default:
			zapLog.Warn("日志通道已满，日志被丢弃")
		}
	}
}

// 缓存处理
func InitOperationLogWorker(zapLog *zap.Logger, db *gorm.DB) {
	go func() {
		for log := range operationLogChan {
			if err := db.Create(&log).Error; err != nil {
				zapLog.Error("日志写入失败:" + err.Error())
			}
		}
	}()
}
