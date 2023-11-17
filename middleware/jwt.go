package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"loopy-manager/global/model"
	"loopy-manager/utils"
	"net/http"
	"time"
)

type CustomClaims struct {
	User model.User `json:"user"`
	jwt.StandardClaims
}

var SigningKey = []byte("Cyw-github")

func CreateToken(user model.User) (string, error) {
	//获取token，前两部分
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{User: user,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),           //签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*2, //2小时过期
			Issuer:    "cyw",                       //签发人，
		},
	})
	//根据密钥生成加密token，token完整三部分
	tokenString, err := token.SignedString(SigningKey)
	if err != nil {
		return "", err
	}
	//存入redis
	err = utils.Redis{}.SetValue(tokenString, user.Account, 60*60*24*time.Second)
	return tokenString, err
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		// 对token对象中的Claim进行类型断言
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
			return claims, nil
		}
		return nil, errors.New("token无效")
	}
	return nil, errors.New("token无效")
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取token
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, utils.Response{Code: 401, Message: "token缺失", Data: ""})
			//终止
			c.Abort()
			return
		}
		claims, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, utils.Response{Code: 401, Message: "token过期", Data: err.Error()})
			//终止
			c.Abort()
			return
		}
		//将用户信息储存再上下文
		c.Set("user", claims.User)
		//重新存入redis
		err = utils.Redis{}.SetValue(token, claims.User.Name, 60*60*24)
		if err != nil {
			fmt.Println(err.Error())
		}
		//继续下面的操作
		c.Next()
	}
}
