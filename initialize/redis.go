package initialize

import (
	"LoopyTicker/global"
	"fmt"
	"github.com/go-redis/redis"
)

// 开启RedisPool
func RedisInit() {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 20,
	})
	ping, err := global.RedisClient.Ping().Result()
	if err != nil {
		fmt.Println("redis连接失败", ping, err)
		return
	}
	fmt.Println("redis连接成功", ping)
}
