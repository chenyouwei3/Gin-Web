package consts

import "time"

const (
	AccessTokenExpireDuration  = 24 * time.Hour      //accessToken刷新时间
	RefreshTokenExpireDuration = 10 * 24 * time.Hour //refreshToken刷新时间
)
