package cacheRedis

import (
	"fmt"
	"time"
)

func (r RedisCache) SetValue(key, value string, t time.Duration) error {
	err := RedisClient.Set(key, value, t).Err()
	if err != nil {
		return fmt.Errorf("redisUtils(sds)设置失败:%w", err)
	}
	return nil
}

func (r RedisCache) GetValue(key string) (string, error) {
	res, err := RedisClient.Get(key).Result()
	if err != nil {
		return "", fmt.Errorf("redisUtils(sds)读取失败:%w", err)
	}
	return res, nil
}

func (r RedisCache) DeletedValue(key string) error {
	err := RedisClient.Del(key).Err()
	if err != nil {
		return fmt.Errorf("redisUtils(sds)删除失败:%w", err)
	}
	return nil
}
