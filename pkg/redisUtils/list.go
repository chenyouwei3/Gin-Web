package redisUtils

import (
	"fmt"
	"loopy-manager/app/global"
)

func (r Redis) SetValueList(key, value string) error {
	err := global.RedisClient.LPush(key, value).Err()
	if err != nil {
		return fmt.Errorf("redisUtils(sds)设置失败:%w", err)
	}
	return nil
}
