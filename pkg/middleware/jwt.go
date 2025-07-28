package middleware

import (
	"gin-web/pkg/extendController"
	"gin-web/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, refreshToken := c.GetHeader("access_token"), c.GetHeader("refresh_token")
		if accessToken == "" {
			c.JSON(200, extendController.Response{
				Code: 4010,
				Message: extendController.ResponseMsg{
					"token为空",
					"token is null",
				},
				Data: "鉴权失败",
			})
			c.Abort()
			return
		}
		//验证token
		newAccessToken, newRefreshToken, err := jwt.ParseRefreshToken(accessToken, refreshToken)
		if err != nil {
			c.JSON(200, extendController.Response{
				Code: 4010,
				Message: extendController.ResponseMsg{
					"token验证失败",
					"Token verification failed",
				},
				Data: "鉴权失败",
			})
			c.Abort()
			return
		}
		SetToken(c, newAccessToken, newRefreshToken)
		c.Next()
	}
}

func SetToken(c *gin.Context, accessToken, refreshToken string) {
	secure := IsHttps(c)
	c.Header("access_token", accessToken)
	c.Header("refresh_token", refreshToken)
	c.SetCookie("access_token", accessToken, 3600*24, "/", "", secure, true)
	c.SetCookie("refresh_token", refreshToken, 3600*24, "/", "", secure, true)
}

// 判断是否https
func IsHttps(c *gin.Context) bool {
	if c.GetHeader("X-Forwarded-Proto") == "https" || c.Request.TLS != nil {
		return true
	}
	return false
}
