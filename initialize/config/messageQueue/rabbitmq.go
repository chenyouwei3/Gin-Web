package messageQueue

import (
	"loopy-manager/initialize/config/messageQueue/rabbitmqUtils"
)

var (
	RabbitCache *rabbitmqUtils.RabbitMQ
)

func RabbitmqInit() {
	RabbitCache = rabbitmqUtils.NewRabbitMqUrl("redisCache", "", "")
	RabbitCache.ConsumeSimple()
}
