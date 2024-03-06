package utils

import (
	"fmt"
	"loopy-manager/initialize/global"
	"time"
)

type Redis struct {
}

func (redis Redis) SetValue(key, value string, t time.Duration) error {
	err := global.RedisClient.Set(key, value, t).Err()
	if err != nil {
		return fmt.Errorf("redis(sds)设置失败:%w", err)
	}
	return nil
}

func (redis Redis) GetValue(key string) (string, error) {
	res, err := global.RedisClient.Get(key).Result()
	if err != nil {
		return "", fmt.Errorf("redis(sds)读取失败:%w", err)
	}
	return res, nil
}
