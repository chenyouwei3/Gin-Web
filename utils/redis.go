package utils

import (
	"loopy-manager/global"
	"time"
)

type Redis struct {
}

func (redis Redis) SetValue(key, value string, t time.Duration) error {
	err := global.RedisClient.Set(key, value, t).Err()
	if err != nil {
		return err
	}
	return nil
}

func (redis Redis) GetValue(key string) (string, error) {
	res, err := global.RedisClient.Get(key).Result()
	if err != nil {
		return "", nil
	}
	return res, nil
}
