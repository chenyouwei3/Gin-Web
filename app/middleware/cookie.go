package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"loopy-manager/pkg/auth"
	"strings"
)

func AuthCookieMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieString := c.Request.Header.Get("Cookie")
		if cookieString == "" {
			c.Abort()
			c.JSON(401, gin.H{
				"code": 1,
				"msg":  "Not authenticate",
			})
			return
		}
		parts := strings.SplitN(cookieString, "=", 2) // 分割字符串，限制分割次数为 2
		key := parts[0]                               // 获取分割后的第一个子串
		value := parts[1]                             // 获取分割后的第二个子串
		cookie, err := c.Request.Cookie(key)
		fmt.Println(auth.CookieDecrypt(key, value), "111111")
		if err != nil {
			c.Abort()
			c.JSON(401, gin.H{
				"code": 1,
				"msg":  "Not authenticate",
			})
			return
		}
		c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		c.Next()
	}
}
