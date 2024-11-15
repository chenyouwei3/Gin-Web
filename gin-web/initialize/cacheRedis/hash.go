package cacheRedis

import "fmt"

func (r RedisCache) SetValueHash(key, field, value string) error {
	err := RedisClient.HSet(key, field, value).Err()
	if err != nil {
		return fmt.Errorf("redisUtils(sds)设置失败:%w", err)
	}
	return nil
}

func (r RedisCache) GetValueHash(key, field string) (string, error) {
	value, err := RedisClient.HGet(key, field).Result()
	if err != nil {
		return "", fmt.Errorf("redisUtils(sds)读取失败:%w", err)
	}
	return value, nil
}

func (r RedisCache) DeleteValueHash(key, field string) error {
	err := RedisClient.HDel(key, field).Err()
	if err != nil {
		return fmt.Errorf("redisUtils(hash)删除失败:%w", err)
	}
	return nil
}
