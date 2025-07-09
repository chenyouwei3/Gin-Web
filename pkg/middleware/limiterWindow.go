package middleware

import (
	"gin-web/pkg/extendController"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 滑动窗口
var LimitQueue map[string][]int64
var ok bool

// 流是限制单个用户在特定时间窗口内的请求数量
// timeWindow(时间窗口) count(请求次数限制)
func LimiterWindow(timeWindow int64, count uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "limit:" + ip //可以根据自己的请求更改
		if !LimitFreqSingle(key, count, timeWindow) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, extendController.Response{
				Code: http.StatusTooManyRequests,
				Message: extendController.ResponseMsg{
					"当前服务器过载,请稍后重试",
					"The current server is overloaded, please try again later",
				},
				Data: nil,
			})
			return
		}
		c.Next()
	}
}

func LimitFreqSingle(queueName string, count uint, timeWindow int64) bool {
	//初始化
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
