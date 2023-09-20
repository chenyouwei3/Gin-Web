package middleware

import (
	"LoopyTicker/model"
	"LoopyTicker/utils"
	"errors"
	"github.com/golang-jwt/jwt"
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

func ResolvingToken(tokenString string) (*CustomClaims, error) {
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
