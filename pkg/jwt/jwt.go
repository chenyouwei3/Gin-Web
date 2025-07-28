package jwt

import (
	"errors"
	"gin-web/consts"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 定义签名密钥，用于加密 JWT
var jwtSecret = []byte("FanOne")

type Claims struct {
	Name               string `json:"name"`
	jwt.StandardClaims        // 嵌入标准声明，包含 JWT 的基本声明
}

// GenerateToken 签发用户Token
func GenerateToken(name string) (accessToken, refreshToken string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(consts.AccessTokenExpireDuration)
	rtExpireTime := nowTime.Add(consts.RefreshTokenExpireDuration)

	claims := Claims{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-web",
		},
	}
	//因为claim包含jwt.StandardClaims 他自己实现了方法，所有可以传入
	// 加密并获得完整的编码后的字符串token
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: rtExpireTime.Unix(),
		Issuer:    "gin-web",
	}).SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// ParseRefreshToken 验证用户token
func ParseRefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	accessClaim, err := ParseToken(aToken)
	if err != nil {
		return
	}
	refreshClaim, err := ParseToken(rToken)
	if err != nil {
		return
	}
	if accessClaim.ExpiresAt > time.Now().Unix() {
		// 如果 access_token 没过期,每一次请求都刷新 refresh_token 和 access_token
		return GenerateToken(accessClaim.Name)
	}

	if refreshClaim.ExpiresAt > time.Now().Unix() {
		// 如果 access_token 过期了,但是 refresh_token 没过期, 刷新 refresh_token 和 access_token
		return GenerateToken(accessClaim.Name)
	}
	// 如果两者都过期了,重新登陆
	return "", "", errors.New("身份过期，重新登陆")
}
