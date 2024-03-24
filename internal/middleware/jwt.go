package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"loopy-manager/pkg/auth/jwt"
	"loopy-manager/pkg/redis"
	"loopy-manager/pkg/utils"
	"net/http"
)

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取token
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, utils.Response{Code: 401, Message: "token缺失", Data: ""})
			//终止
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, utils.Response{Code: 401, Message: "token过期", Data: err.Error()})
			//终止
			c.Abort()
			return
		}
		//将用户信息储存再上下文
		c.Set("user", claims.User)
		//重新存入redis
		err = redis.Redis{}.SetValue(token, claims.User.Name, 60*60*48)
		if err != nil {
			logrus.Error("更新token失败:", err)
		}
		//继续下面的操作
		c.Next()
	}
}
