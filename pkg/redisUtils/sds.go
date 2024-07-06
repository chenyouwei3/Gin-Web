package redisUtils

import (
	"fmt"
	"loopy-manager/app/global"
	"time"
)

type Redis struct {
}

func (r Redis) SetValue(key, value string, t time.Duration) error {
	err := global.RedisClient.Set(key, value, t).Err()
	if err != nil {
		return fmt.Errorf("redisUtils(sds)设置失败:%w", err)
	}
	return nil
}

func (r Redis) GetValue(key string) (string, error) {
	res, err := global.RedisClient.Get(key).Result()
	if err != nil {
		return "", fmt.Errorf("redisUtils(sds)读取失败:%w", err)
	}
	return res, nil
}

func (r Redis) DeletedValue(key string) error {
	err := global.RedisClient.Del(key).Err()
	if err != nil {
		return fmt.Errorf("redisUtils(sds)删除失败:%w", err)
	}
	return nil
}
