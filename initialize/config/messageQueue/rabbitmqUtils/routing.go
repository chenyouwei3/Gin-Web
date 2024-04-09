package rabbitmqUtils

import (
	"github.com/streadway/amqp"
)

func (r *RabbitMQ) PublishRouting(message string) {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	err := r.Channel.ExchangeDeclare(
		r.ExChange, //交换机名字
		"direct",   //交换机类型
		true,       //是否持久化
		false,      //是否自动删除，表示交换机在没有与之绑定的队列时是否自动删除
		false,      // 是否是内部交换机，表示该交换机是否被客户端使用
		false,      // 是否等待服务器响应，这里设置为 false。
		nil,
	)
	r.failOnErr("创建交换机失败", err)
	err = r.Channel.Publish(
		r.ExChange, // 表示将消息发送到默认的交换机
		r.Key,      // 指定了消息的路由键，即将消息发送到的队列
		false,      //是 mandatory 参数，它表示如果无法将消息路由到队列中， 消息会被丢弃而不是返回给发送者。
		false,      //是 immediate 参数，它表示如果无法立即将消息发送给接收者， 消息会被丢弃而不是返回给发送者
		amqp.Publishing{
			//DeliveryMode: amqp.Persistent, //是否持久化
			ContentType: "text/plain",    //ContentType 表示消息的类型为 “text/plain”，
			Body:        []byte(message), //Body 则是将消息内容转化为字节流。
		})
	r.failOnErr("推送信息通道错误:", err)
}

func (r *RabbitMQ) ConsumeRouting(Func func(message string)) {
	err := r.Channel.ExchangeDeclare(
		r.ExChange, //交换机名字
		"direct",   //交换机类型
		true,       //是否持久化
		false,      //是否自动删除，表示交换机在没有与之绑定的队列时是否自动删除
		false,      // 是否是内部交换机，表示该交换机是否被客户端使用
		false,      // 是否等待服务器响应，这里设置为 false。
		nil,
	)
	r.failOnErr("创建交换机失败", err)
	q, err := r.Channel.QueueDeclare(
		"",    // name
		false, // 持久的
		false, //当队列不再使用时是否自动删除
		true,  //是否设置排他性队列（只能由声明它的连接使用）
		false, //是否设置为无等待（no-wait，等待服务器响应）。
		nil,
	)
	r.failOnErr("声明发送的消息队列失败", err)
	//获取接受消息的Delivery通道
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
	err = r.Channel.QueueBind(
		q.Name,     // queue name
		r.Key,      // 路由键
		r.ExChange, // exchange
		false,      //指示是否等待服务器的响应。在这里，设置为 false 表示要等待服务器的响应
		nil,
	)
	r.failOnErr("绑定失败", err)
	go func() {
		for d := range msgs {
			Func(string(d.Body))
		}
	}()
	select {}
}
