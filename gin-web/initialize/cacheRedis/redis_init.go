package cacheRedis

import (
	conf "gin-web/initialize/config"
	"github.com/go-redis/redis"
	"log"
)

var RedisClient *redis.Client

type RedisCache struct{}

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         conf.Conf.Redis.Addr + ":" + conf.Conf.Redis.Port, // Redis 服务器地址
		Password:     conf.Conf.Redis.Password,                          // Redis 服务器密码
		DB:           conf.Conf.Redis.DB,                                // Redis 数据库索引
		PoolSize:     conf.Conf.Redis.PoolSize,                          // 连接池大小
		MinIdleConns: conf.Conf.Redis.MinIdleConns,                      // 最小空闲连接数
	})
	ping, err := RedisClient.Ping().Result()
	if err != nil {
		return err
	}
	log.Println("Redis connection successful", ping)
	return nil
}
