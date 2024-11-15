package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 滑动窗口
var LimitQueue map[string][]int64
var ok bool

func LimiterWindow(c *gin.Context, timeWindow int64, count uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "limit:" + ip
		if !LimitFreqSingle(key, count, timeWindow) {
			c.JSON(200, gin.H{
				"code": http.StatusTooManyRequests,
				"msg":  "当前服务器过载,请稍后重试",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func LimitFreqSingle(queueName string, count uint, timeWindow int64) bool {
	currTime := time.Now().Unix() //获取当前时间
	if LimitQueue == nil {        //检查全局变量 初始化
		LimitQueue = make(map[string][]int64)
	}
	if _, ok := LimitQueue[queueName]; !ok { //没东西的话
		LimitQueue[queueName] = make([]int64, 0)
	}
	//小于设定的最大数
	if uint(len(LimitQueue[queueName])) < count {
		LimitQueue[queueName] = append(LimitQueue[queueName], currTime)
		return true
	}
	//大于设定的最大数
	earlyTime := LimitQueue[queueName][0]
	//最早的请求还在时间窗口内，表示请求过于频繁
	if currTime-earlyTime <= timeWindow {
		return false
	}
	//最早的请求不在时间窗口内，表示请求过于频繁
	LimitQueue[queueName] = LimitQueue[queueName][1:] //去掉 LimitQueue[queueName] 的第一个元素
	LimitQueue[queueName] = append(LimitQueue[queueName], currTime)
	return true
}
