package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

func Limiter(r rate.Limit, b int) gin.HandlerFunc {
	//第一个参数是r Limit，这是代表每秒可以向令牌桶中产生多少令牌
	//第二个参数是b int，这是代表令牌桶的容量大小
	limiter := rate.NewLimiter(r, b)
	return func(c *gin.Context) {
		//AllowN 方法表示，截止到某一时刻，
		//目前桶中数目是否至少为 n 个，满足则返回 true，同时从桶中消费 n 个 token。
		if !limiter.Allow() {
			c.String(http.StatusTooManyRequests, "Too Many Requests")
			c.Abort()
			return
		}
		c.Next()
	}
}
