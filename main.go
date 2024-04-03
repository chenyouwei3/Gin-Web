package main

import (
	"fmt"
	"log"
	"loopy-manager/initialize"
	"loopy-manager/initialize/global"
)

func init() {
	initialize.InitConfig()
}

func main() {
	//engine := router.GetEngine()
	//if err := engine.Run(":8098"); err != nil {
	//	panic(err)
	//}
	//global.RedisClient.Set("user:1", "data1", 0)
	//global.RedisClient.Set("user:2", "data2", 0)
	//global.RedisClient.Set("role:1", "data3", 0)
	//global.RedisClient.Set("role:2", "data4", 0)
	//global.RedisClient.Set("api:1", "data5", 0)
	//global.RedisClient.Set("api:2", "data6", 0)

	// 删除以 "user:" 开头的所有缓存键
	keys, _ := global.RedisClient.Keys("api:*").Result()
	// 删除符合条件的所有缓存键
	fmt.Println(keys)
	for _, key := range keys {
		if err := global.RedisClient.Del(key).Err(); err != nil {
			log.Fatalf("删除 Redis 缓存键 %s 失败：%v", key, err)
		} else {
			fmt.Printf("删除 Redis 缓存键 %s 成功\n", key)
		}
	}
}
