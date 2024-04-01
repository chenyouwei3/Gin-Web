package messageQueue

import (
	"loopy-manager/initialize/global"
	"loopy-manager/pkg/rabbitmqUtils"
)

func RabbitmqInit() {
	global.RabbitCache = rabbitmqUtils.NewRabbitMqUrl("redisCache", "", "")
}
