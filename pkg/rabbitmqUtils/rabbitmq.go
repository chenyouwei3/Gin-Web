package rabbitmqUtils

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"loopy-manager/initialize/config"
)

//交换机种类direct topic headers fanout

type RabbitMQ struct {
	Conn      *amqp.Connection
	Channel   *amqp.Channel
	QueueName string //队列名称
	ExChange  string //交换机名称
	Key       string //消息的路由键，即将消息发送到的队列
	Url       string //连接地址
}

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, ExChange: exchange, Key: key, Url: config.Config.Rabbitmq.Url1}
}

// Destroy 断开channel 和 connection
func (r *RabbitMQ) Destroy() {
	r.Channel.Close()
	r.Conn.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(message string, err error) {
	if err != nil {
		logrus.Error(message, ":", err)
		return
	}
}

// 删除队列
func (r *RabbitMQ) QueueDelete(QueueName string) {
	_, err := r.Channel.QueueDelete(QueueName, false, false, false)
	r.failOnErr("删除队列失败:", err)
}
