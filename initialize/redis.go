package initialize

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"loopy-manager/global"
)

func RedisInit() {
	if global.RedisPool == nil {
		global.RedisPool = &redis.Pool{
			MaxIdle:     3, //最大空闲连接数
			MaxActive:   0, //最大链接数，0无限制
			IdleTimeout: 0, //连接超时时间不限制
			Dial: func() (redis.Conn, error) {
				dial, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("password"))
				if err != nil {
					log.Fatalln("redis连接失败", err.Error())
				}
				return dial, err
			},
		}
		global.RedisPool.Get()
		fmt.Println("redis连接成功！")
	}
}
