package middleware

import (
	"github.com/gin-gonic/gin"
	"loopy-manager/global/model"
	"time"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path      //请求路径
		query := c.Request.URL.RawQuery //query参数
		c.Next()

		var body interface{}

		cost := time.Since(start) //访问时间
		user, _ := c.Get("user")
		log := model.Log{
			User:      user,
			Path:      path,
			Method:    c.Request.Method,
			Status:    c.Writer.Status(),
			Query:     query,
			Body:      body,
			Ip:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Errors:    c.Errors.ByType(gin.ErrorTypePrivate).String(),
			Cost:      cost.String(),
		}
		if log.Status == 204 {
			return
		}

	}
}
