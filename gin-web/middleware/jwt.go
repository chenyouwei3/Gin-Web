package middleware

import (
	"errors"
	"gin-web/initialize/cacheRedis"
	"gin-web/initialize/runLog"
	"gin-web/utils/extendController"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取token
		token := c.Request.Header.Get("token")
		if token == "" {
			runLog.ZapLog.Info("Identity information not passed")
			extendController.BaseController{}.SendUnAuthResponse(c)
			//终止
			c.Abort()
			return
		}
		claims, err := ParseToken(token)
		if err != nil {
			runLog.ZapLog.Info("Identity information not passed")
			extendController.BaseController{}.SendUnAuthResponse(c)
			//终止
			c.Abort()
			return
		}
		//将用户信息储存再上下文
		//c.Set("user", claims.Issuer)
		//重新存入redis
		err = cacheRedis.RedisCache{}.SetValue(token, claims.Issuer, 60*60*48)
		if err != nil {
			runLog.ZapLog.Info(err.Error())
		}
		//继续下面的操作
		c.Next()
	}
}

// 定义一个自定义声明结构体
type CustomClaims struct {
	jwt.StandardClaims // 嵌入标准声明，包含 JWT 的基本声明
}

// 定义签名密钥，用于加密 JWT
var SigningKey = "SmartGraphiteBySDL"

// 创建 JWT 的函数，接受一个用户对象作为参数
func GenerateToken(issuer string) (string, error) {
	// 获取 token，指定使用 HS256 签名方法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(), // 签名生效时间（当前时间）
			// ExpiresAt: time.Now().Unix() + 60*60*24*30, // 可选：设置过期时间，30 天后过期
			Issuer: issuer, // 签发人
		},
	})
	// 生成 JWT 字符串，包含头部、载荷和签名
	tokenString, err := token.SignedString([]byte(SigningKey)) // 使用签名密钥对 token 进行签名
	return tokenString, err                                    // 返回生成的 token
}

func ParseToken(tokenString string) (CustomClaims, error) {
	//解析token
	var claims CustomClaims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SigningKey), nil
	})
	if err != nil {
		return CustomClaims{}, err
	}
	//再redis中查看token是否过期
	_, err = cacheRedis.RedisCache{}.GetValue(tokenString)
	if err != nil {
		return CustomClaims{}, errors.New("token过期")
	}
	return claims, err
}
