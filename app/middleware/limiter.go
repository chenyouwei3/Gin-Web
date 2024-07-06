package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func LimiterBucket(r rate.Limit, b int) gin.HandlerFunc {
	//第一个参数是r Limit，这是代表每秒可以向令牌桶中产生多少令牌
	//第二个参数是b int，这是代表令牌桶的容量大小
	limiter := rate.NewLimiter(r, b)
	return func(c *gin.Context) {
		//AllowN 方法表示，截止到某一时刻，
		//目前桶中数目是否至少为 n 个，满足则返回 true，同时从桶中消费 n 个 token。
		if !limiter.Allow() {
			c.String(http.StatusTooManyRequests, "当前服务器过载,请稍后重试")
			c.Abort()
			return
		}
		c.Next()
	}
}

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

var LimitQueue map[string][]int64
var ok bool

func LimitFreqSingle(queueName string, count uint, timeWindow int64) bool {
	currTime := time.Now().Unix()
	if LimitQueue == nil {
		LimitQueue = make(map[string][]int64)
	}
	if _, ok = LimitQueue[queueName]; !ok {
		LimitQueue[queueName] = make([]int64, 0)
	}
	//队列未满
	if uint(len(LimitQueue[queueName])) < count {
		LimitQueue[queueName] = append(LimitQueue[queueName], currTime)
		return true
	} //队列满了，取出最早访问的时间
	earlyTime := LimitQueue[queueName][0]
	//说明最早期的时间还在时间窗口内,还没过期,所以不允许通过
	if currTime-earlyTime <= timeWindow {
		return false
	} else {
		//说明最早期的访问应该过期了,去掉最早期的
		LimitQueue[queueName] = LimitQueue[queueName][1:]
		LimitQueue[queueName] = append(LimitQueue[queueName], currTime)
	}
	return true
}
