package middleware

import (
	"github.com/golang-jwt/jwt"
	"loopy-manager/model"
	"time"
)

type CustomClaims struct {
	User model.User `json:"user"`
	jwt.StandardClaims
}

var SigningKey = []byte("Chenyouwei3")

func CreateToken(user model.User) (string, error) {
	claims := CustomClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(), //签名生效时间
			ExpiresAt: time.Now().Unix(), //有效时间
			Issuer:    "cyw",             //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	//根据密钥生成加密token，token完整三部分
	//获取完整的名牌
	tokenString, err := token.SignedString(SigningKey)
	if err != nil {
		return "", err
	}
	//存入redis
	return tokenString, err
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	56565
}
