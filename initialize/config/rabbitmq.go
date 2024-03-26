package config

import (
	"loopy-manager/initialize/global"
	"loopy-manager/pkg/rabbitmqUtils"
	"loopy-manager/pkg/redisUtils"
)

func RabbitmqInit() {
	global.RabbitCache = rabbitmqUtils.NewRabbitMQ("redisCache", "", "")
	RabbitmqConsume()
}

func RabbitmqConsume() {
	global.RabbitCache.ConsumeSimple(redisUtils.Redis{}.DeletedValue)
}
