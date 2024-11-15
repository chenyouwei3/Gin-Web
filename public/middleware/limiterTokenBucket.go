package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

// 令牌桶
func LimiterTokenBucket(r rate.Limit, b int) gin.HandlerFunc {
	//r 每秒生成令牌数量,即请求速率
	limiter := rate.NewLimiter(r, b)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.String(http.StatusTooManyRequests, "当前服务器过载,请稍后重试")
			c.Abort()
			return
		}
		c.Next()
	}
}
