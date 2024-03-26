package rabbitmqUtils

import (
	"github.com/streadway/amqp"
)

func (r *RabbitMQ) PublishExchange(message string) {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	err := r.Channel.ExchangeDeclare(
		r.ExChange, //交换机名字
		"fanout",   //交换机类型
		true,       //是否持久化
		false,      //是否自动删除，表示交换机在没有与之绑定的队列时是否自动删除
		false,      // 是否是内部交换机，表示该交换机是否被客户端使用
		false,      // 是否等待服务器响应，这里设置为 false。
		nil,
	)
	r.failOnErr("创建交换机失败", err)
	//调用channel 发送消息到队列中
	r.Channel.Publish(
		r.ExChange, // 表示将消息发送到默认的交换机
		r.Key,      // 指定了消息的路由键，即将消息发送到的队列
		false,      //是 mandatory 参数，它表示如果无法将消息路由到队列中， 消息会被丢弃而不是返回给发送者。
		false,      //是 immediate 参数，它表示如果无法立即将消息发送给接收者， 消息会被丢弃而不是返回给发送者
		amqp.Publishing{
			//DeliveryMode: amqp.Persistent, //是否持久化
			ContentType: "text/plain",    //ContentType 表示消息的类型为 “text/plain”，
			Body:        []byte(message), //Body 则是将消息内容转化为字节流。
		})
}

func (r *RabbitMQ) ConsumeExchange(Func func(message string)) {
	err := r.Channel.ExchangeDeclare(
		r.ExChange, //交换机名字
		"fanout",   //交换机类型
		true,       //是否持久化
		false,      //是否自动删除，表示交换机在没有与之绑定的队列时是否自动删除
		false,      // 是否是内部交换机，表示该交换机是否被客户端使用
		false,      // 是否等待服务器响应，这里设置为 false。
		nil,
	)
	r.failOnErr("创建交换机失败", err)
	q, err := r.Channel.QueueDeclare(
		"",    //
		false, // 队列是否持久化（在 RabbitMQ 服务重启后仍然存在）
		false, // 当队列不再使用时是否自动删除
		true,  //是否设置排他性队列（只能由声明它的连接使用）
		false, //是否设置为无等待（no-wait，等待服务器响应）。
		nil,
	)
	r.failOnErr("声明消息队列失败", err)
	err = r.Channel.QueueBind(
		q.Name, // queue name
		r.Key,  // 路由键
		"logs", // exchange
		false,  //指示是否等待服务器的响应。在这里，设置为 false 表示要等待服务器的响应
		nil,
	)
	r.failOnErr("绑定失败", err)
	msgs, err := r.Channel.Consume(
		q.Name, // 队列名称，表示要注册消费者的队列的名称。
		"",     // 消费者名称，表示要为消费者指定一个名称。在这里，使用一个空字符串表示使用默认的消费者名称。
		true,   // auto-ack 参数，指示消费者是否自动确认收到的消息。在这里，设置为 true 表示消息被消费者接收后会自动发送确认给服务器。
		false,  // exclusive 参数，指示队列是否是独占的。在这里，设置为 false 表示队列可以被其他连接访问。
		false,  // no-local 参数，指示是否发送 no-local 标志给服务器，如果设置为 true，则不能将消息发送给自己。在这里，设置为 false 表示可以发送消息给自己。
		false,  // no-wait 参数，指示是否等待服务器的响应。在这里，设置为 false 表示要等待服务器的响应。
		nil,
	)
	r.failOnErr("Failed to register a consumer", err)
	go func() {
		for d := range msgs {
			Func(string(d.Body))
		}
	}()
	select {}
}
