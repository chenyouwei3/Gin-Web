package redis

import (
	"fmt"
	"loopy-manager/initialize/global"
)

func (r Redis) SetValueHash(key, field, value string) error {
	err := global.RedisClient.HSet(key, field, value).Err()
	if err != nil {
		return fmt.Errorf("redis(sds)设置失败:%w", err)
	}
	return nil
}

func (r Redis) GetValueHash(key, field string) (string, error) {
	value, err := global.RedisClient.HGet(key, field).Result()
	if err != nil {
		return "", fmt.Errorf("redis(sds)读取失败:%w", err)
	}
	return value, nil
}

func (r Redis) DeleteValueHash(key, field string) error {
	err := global.RedisClient.HDel(key, field).Err()
	if err != nil {
		return fmt.Errorf("redis(hash)删除失败:%w", err)
	}
	return nil
}
