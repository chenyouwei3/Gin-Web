package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"loopy-manager/app/model"
	"loopy-manager/initialize/config"
	"loopy-manager/pkg/redisUtils"
	"time"
)

type CustomClaims struct {
	User model.User `json:"user"`
	jwt.StandardClaims
}

var SigningKey = []byte(config.Config.Jwt.SignKey)

func CreateToken(user model.User) (string, error) {
	//获取token，前两部分
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{User: user,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               //签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //30添小时过期
			Issuer:    config.Config.Jwt.Issuer,        //签发人，
		},
	})
	//根据密钥生成加密token，token完整三部分
	tokenString, err := token.SignedString(SigningKey)
	if err != nil {
		return "", err
	}
	//存入redis
	err = redisUtils.Redis{}.SetValue(tokenString, user.Name, 60*60*168)
	if err != nil {
		return "", err
	}
	return tokenString, err
}
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	//再redis中查看token是否过期
	_, err = redisUtils.Redis{}.GetValue(tokenString)
	if err != nil {
		return nil, errors.New("token过期")
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("token无效")
	}
	return nil, errors.New("token无效")
}
