package middleware

import (
	"gin-web/utils/extendController"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

// 令牌桶
// r[令牌生成速率(即每秒产生多少个令牌)] b[令牌桶的最大容量(即最多存多少个令牌)]
func LimiterTokenBucket(r rate.Limit, b int) gin.HandlerFunc {
	//r 每秒生成令牌数量,即请求速率
	limiter := rate.NewLimiter(r, b)
	return func(c *gin.Context) {
		if !limiter.Allow() {
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
