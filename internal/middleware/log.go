package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"loopy-manager/initialize/global"
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
		var user struct {
			account  string `json:"account"`
			password string `json:"password"`
		}
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			logrus.Error("日志结构体解析body失败1:", err)
		}
		//fmt.Println("body", string(body))
		err = json.Unmarshal(body, &user)
		if err != nil {
			logrus.Error("日志结构体绑定body失败2", err)
		}
		operationLog := OperationLog{
			Id:        global.LogSnowFlake.Generate().Int64(),
			Username:  user.account,
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
		go func(log OperationLog) {
			res := global.LogTable.Create(log)
			if res.Error != nil {
				logrus.Error("中间件日志记录失败:", res.Error)
			}
		}(operationLog)
		c.Next()
	}
}

type OperationLog struct {
	Id        int64     `json:"id" gorm:"column:id;type:bigint;primarykey;not null"`
	Username  string    `gorm:"type:varchar(20);comment:'用户登录名'" json:"username"`
	Ip        string    `gorm:"type:varchar(20);comment:'Ip地址'" json:"ip"`
	Method    string    `gorm:"type:varchar(20);comment:'请求方式'" json:"method"`
	Query     string    `gorm:"type:varchar(50)" json:"query"`
	Path      string    `gorm:"type:varchar(100);comment:'访问路径'" json:"path"`
	Status    int       `gorm:"type:int(4);comment:'响应状态码'" json:"status"`
	StartTime time.Time `gorm:"type:datetime(3);comment:'发起时间'" json:"startTime"`
	TimeCost  int64     `gorm:"type:int(6);comment:'请求耗时(ms)'" json:"timeCost"`
	UserAgent string    `gorm:"type:varchar(50);comment:'浏览器标识'" json:"userAgent"`
	Errors    string    `gorm:"type:varchar(100)"json:"errors"`
}
