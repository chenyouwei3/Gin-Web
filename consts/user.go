package consts

import "time"

const (
	AccessTokenExpireDuration  = 2 * time.Hour      //accessToken刷新时间
	RefreshTokenExpireDuration = 2 * 12 * time.Hour //refreshToken刷新时间
)
