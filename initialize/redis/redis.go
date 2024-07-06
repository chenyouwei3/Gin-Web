package redis

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"loopy-manager/app/global"
	"loopy-manager/initialize/config"
)

// RedisInit 开启RedisPool
func RedisInit(config config.RedisConfig) {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:         config.Address,      // Redis 服务器地址
		Password:     config.Password,     // Redis 服务器密码
		DB:           config.DB,           // Redis 数据库索引
		PoolSize:     config.PoolSize,     // 连接池大小
		MinIdleConns: config.MinIdleConns, // 最小空闲连接数
	})
	ping, err := global.RedisClient.Ping().Result()
	if err != nil {
		logrus.Fatalln("redis连接失败", ping, err)
	}
	logrus.Println("redis连接成功", ping)
}
