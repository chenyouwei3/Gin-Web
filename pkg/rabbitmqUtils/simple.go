package rabbitmqUtils

import (
	"fmt"
	"github.com/streadway/amqp"
)

func (r *RabbitMQ) PublishSimple(message string) {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err := r.Channel.QueueDeclare(
		r.QueueName,
		false, // 持久的
		false, //当队列不再使用时是否自动删除
		false, //是否设置排他性队列（只能由声明它的连接使用）
		false, //是否设置为无等待（no-wait，等待服务器响应）。
		nil,
	)
	r.failOnErr("声明发送的消息队列失败", err)
	//调用channel 发送消息到队列中
	r.Channel.Publish(
		r.ExChange,
		r.QueueName,
		false, //是 mandatory 参数，它表示如果无法将消息路由到队列中，消息会被丢弃而不是返回给发送者。
		false, //是 immediate 参数，它表示如果无法立即将消息发送给接收者，消息会被丢弃而不是返回给发送者,
		amqp.Publishing{
			//DeliveryMode: amqp.Persistent, //是否持久化
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// ConsumeSimple simple 模式下消费者
func (r *RabbitMQ) ConsumeSimple(Func func(message string) error) {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := r.Channel.QueueDeclare(
		r.QueueName,
		false, // 持久的
		false, //当队列不再使用时是否自动删除
		false, //是否设置排他性队列（只能由声明它的连接使用）
		false, //是否设置为无等待（no-wait，等待服务器响应）。
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	r.failOnErr("声明接收的消息队列失败", err)
	//接收消息
	msgs, err := r.Channel.Consume(
		q.Name, //表示要消费的队列名称
		"",     //表示消费者的名称，这里为空字符串。
		true,   //表示开启自动确认模式，即当消费者接收到消息后自动向 RabbitMQ 发送消息已被消费的回执(消息确认)
		false,  //表示不加锁。
		false,  //表示不禁止消费者使用同一连接发送消息。
		false,  //表示不等待接收返回的参数。
		nil,    //表示不传递其他参数。
	)
	r.failOnErr("创建消费者失败", err)
	go func() {
		for d := range msgs {
			Func(string(d.Body))
		}
	}()
	select {}
}
