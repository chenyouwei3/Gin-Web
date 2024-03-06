package utils

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"loopy-manager/initialize/global"
)

type Rabbit struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string //队列名称
	Exchange  string //交换机名称
	Key       string //bind Key 名称
	RabbitUrl string //连接信息
}

func NewRabbit(queueName string, exchange string, key string) *Rabbit {
	return &Rabbit{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		RabbitUrl: global.RabbitMqAddress,
	}
}

func (r *Rabbit) Close() {
	r.channel.Close()
	r.conn.Close()
}

func NewRabbitConn(queueName, exchange, key string) *Rabbit {
	rabbit := NewRabbit(queueName, exchange, key)
	var err error
	rabbit.conn, err = amqp.Dial(rabbit.RabbitUrl)
	rabbit.failOnErr("创建连接失败:", err)
	rabbit.channel, err = rabbit.conn.Channel()
	rabbit.failOnErr("创建通道失败:", err)
	return rabbit
}

func (r *Rabbit) PublishSimple(message string) {
	_, err := r.channel.QueueDeclare( //1.申请队列，如果队列不存在会自动创建，存在则跳过创建
		r.QueueName, //队列名
		false,       //持久化
		false,       //队列不再使用时是否删除
		false,       //是否设置排他性队列（只能由声明它的连接使用）
		false,       //是否设置为无等待（no-wait，等待服务器响应）。
		nil,
	)
	r.failOnErr("声明发送的消息队列失败1:", err)
	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false, //是 mandatory 参数，它表示如果无法将消息路由到队列中，消息会被丢弃而不是返回给发送者。
		false, //是 immediate 参数，它表示如果无法立即将消息发送给接收者，消息会被丢弃而不是返回给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	r.failOnErr("推送消息失败:", err)
}

func (r *Rabbit) PublishWork(message string) {
	_, err := r.channel.QueueDeclare( //1.申请队列，如果队列不存在会自动创建，存在则跳过创建
		r.QueueName, //队列名
		true,        //持久化
		false,       //队列不再使用时是否删除
		false,       //是否设置排他性队列（只能由声明它的连接使用）
		false,       //是否设置为无等待（no-wait，等待服务器响应）。
		nil,
	)
	r.failOnErr("声明发送的消息队列失败2:", err)
	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false, //是 mandatory 参数，它表示如果无法将消息路由到队列中，消息会被丢弃而不是返回给发送者。
		false, //是 immediate 参数，它表示如果无法立即将消息发送给接收者，消息会被丢弃而不是返回给发送者
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, //是否持久化
			ContentType:  "text/plain",
			Body:         []byte(message),
		})
	r.failOnErr("推送消息失败:", err)
}

func (r *Rabbit) PublishPublish(message string) {
	err := r.channel.ExchangeDeclare(
		r.Exchange, //交换机名字
		"fanout",   //交换机类型
		true,       //是否持久化
		false,      //是否自动删除，表示交换机在没有与之绑定的队列时是否自动删除
		false,      // 是否是内部交换机，表示该交换机是否被客户端使用
		false,      // 是否等待服务器响应，这里设置为 false。
		nil,
	)
	r.failOnErr("创建交换机失败:", err)
	err = r.channel.Publish(
		r.Exchange, // 表示将消息发送到默认的交换机
		"",         // 指定了消息的路由键，即将消息发送到的队列
		false,      //是 mandatory 参数，它表示如果无法将消息路由到队列中， 消息会被丢弃而不是返回给发送者。
		false,      //是 immediate 参数，它表示如果无法立即将消息发送给接收者， 消息会被丢弃而不是返回给发送者
		amqp.Publishing{
			ContentType: "text/plain",    //ContentType 表示消息的类型为 “text/plain”，
			Body:        []byte(message), //Body 则是将消息内容转化为字节流。
		})
	r.failOnErr("推送消息失败:", err)
}

func (r *Rabbit) PublishRouting(message string) {
	err := r.channel.ExchangeDeclare(
		r.Exchange, //交换机名字
		"direct",   //交换机类型
		true,       //是否持久化
		false,      //是否自动删除，表示交换机在没有与之绑定的队列时是否自动删除
		false,      // 是否是内部交换机，表示该交换机是否被客户端使用
		false,      // 是否等待服务器响应，这里设置为 false。
		nil,
	)
	r.failOnErr("创建交换机失败:", err)
	err = r.channel.Publish(
		r.Exchange, // 表示将消息发送到默认的交换机
		r.Key,      // 指定了消息的路由键，即将消息发送到的队列
		false,      //是 mandatory 参数，它表示如果无法将消息路由到队列中， 消息会被丢弃而不是返回给发送者。
		false,      //是 immediate 参数，它表示如果无法立即将消息发送给接收者， 消息会被丢弃而不是返回给发送者
		amqp.Publishing{
			//DeliveryMode: amqp.Persistent, //是否持久化
			ContentType: "text/plain",    //ContentType 表示消息的类型为 “text/plain”，
			Body:        []byte(message), //Body 则是将消息内容转化为字节流。
		})
	r.failOnErr("推送消息失败:", err)
}

func (r *Rabbit) PublishTopic(message string) {
	err := r.channel.ExchangeDeclare(
		r.Exchange, //交换机名字
		"topic",    //交换机类型
		true,       //是否持久化
		false,      //是否自动删除，表示交换机在没有与之绑定的队列时是否自动删除
		false,      // 是否是内部交换机，表示该交换机是否被客户端使用
		false,      // 是否等待服务器响应，这里设置为 false。
		nil,
	)
	r.failOnErr("创建交换机失败:", err)
	err = r.channel.Publish(
		r.Exchange, // 表示将消息发送到默认的交换机
		r.Key,      // 指定了消息的路由键，即将消息发送到的队列
		false,      //是 mandatory 参数，它表示如果无法将消息路由到队列中， 消息会被丢弃而不是返回给发送者。
		false,      //是 immediate 参数，它表示如果无法立即将消息发送给接收者， 消息会被丢弃而不是返回给发送者
		amqp.Publishing{
			//DeliveryMode: amqp.Persistent, //是否持久化
			ContentType: "text/plain",    //ContentType 表示消息的类型为 “text/plain”，
			Body:        []byte(message), //Body 则是将消息内容转化为字节流。
		})
	r.failOnErr("推送消息失败:", err)
}

func (r *Rabbit) ConsumeSimple() {
	q, err := r.channel.QueueDeclare( //1.申请队列，如果队列不存在会自动创建，存在则跳过创建
		r.QueueName, //队列名
		false,       //持久化
		false,       //队列不再使用时是否删除
		false,       //是否设置排他性队列（只能由声明它的连接使用）
		false,       //是否设置为无等待（no-wait，等待服务器响应）。
		nil,
	)
	r.failOnErr("声明接受的消息队列失败:", err)
	msgs, err := r.channel.Consume(
		q.Name, //队列名字
		"",     //表示消费者的名称，这里为空字符串。
		true,   //表示开启自动确认模式，即当消费者接收到消息后自动向 RabbitMQ 发送消息已被消费的回执(消息确认)
		false,  //表示不加锁。
		false,  //表示不禁止消费者使用同一连接发送消息。
		false,  //表示不等待接收返回的参数。
		nil,    //表示不传递其他参数。
	)
	r.failOnErr("创建消费者失败:", err)
	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msgs {
			//消息逻辑处理，可以自行设计逻辑
			log.Printf("收到消息: %s", d.Body)

		}
	}()
	<-forever
}

func (r *Rabbit) ConsumeWork() {
	q, err := r.channel.QueueDeclare( //1.申请队列，如果队列不存在会自动创建，存在则跳过创建
		r.QueueName, //队列名
		true,        //持久化
		false,       //队列不再使用时是否删除
		false,       //是否设置排他性队列（只能由声明它的连接使用）
		false,       //是否设置为无等待（no-wait，等待服务器响应）。
		nil,
	)
	r.failOnErr("声明接受的消息队列失败:", err)
	//接收消息
	//*---------------------公平分发------------------*/
	err = r.channel.Qos(
		1,
		0,
		false,
	)
	r.failOnErr("分发失败:", err)
	msgs, err := r.channel.Consume(
		q.Name, //队列名字
		"",     //表示消费者的名称，这里为空字符串。
		false,  //表示开启自动确认模式，即当消费者接收到消息后自动向 RabbitMQ 发送消息已被消费的回执(消息确认)
		false,  //表示不加锁。
		false,  //表示不禁止消费者使用同一连接发送消息。
		false,  //表示不等待接收返回的参数。
		nil,    //表示不传递其他参数。
	)
	r.failOnErr("创建消费者失败:", err)
	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msgs {
			//消息逻辑处理，可以自行设计逻辑
			log.Printf("收到消息: %s", d.Body)

		}
	}()
	<-forever
}

func (r *Rabbit) ConsumePublish() {
	err := r.channel.ExchangeDeclare(
		r.Exchange, //交换机名字
		"fanout",   //交换机类型
		true,       //是否持久化
		false,      //是否自动删除，表示交换机在没有与之绑定的队列时是否自动删除
		false,      // 是否是内部交换机，表示该交换机是否被客户端使用
		false,      // 是否等待服务器响应，这里设置为 false。
		nil,
	)
	r.failOnErr("创建交换机失败:", err)
	q, err := r.channel.QueueDeclare( //1.申请队列，如果队列不存在会自动创建，存在则跳过创建
		"",    //队列名
		false, //持久化
		false, //队列不再使用时是否删除
		true,  //是否设置排他性队列（只能由声明它的连接使用）
		false, //是否设置为无等待（no-wait，等待服务器响应）。
		nil,
	)
	r.failOnErr("声明接受的消息队列失败:", err)
	err = r.channel.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil,
	)
	r.failOnErr("绑定路由失败", err)

	msgs, err := r.channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	r.failOnErr("创建消费者失败:", err)
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()
	<-forever
}

func (r *Rabbit) ConsumeRouting() {
	err := r.channel.ExchangeDeclare(
		r.Exchange, //交换机名字
		"direct",   //交换机类型
		true,       //是否持久化
		false,      //是否自动删除，表示交换机在没有与之绑定的队列时是否自动删除
		false,      // 是否是内部交换机，表示该交换机是否被客户端使用
		false,      // 是否等待服务器响应，这里设置为 false。
		nil,
	)
	r.failOnErr("创建交换机失败", err)

	q, err := r.channel.QueueDeclare(
		"",    // name
		false, // 持久的
		false, //当队列不再使用时是否自动删除
		true,  //是否设置排他性队列（只能由声明它的连接使用）
		false, //是否设置为无等待（no-wait，等待服务器响应）。
		nil,
	)
	r.failOnErr("声明发送的消息队列失败", err)
	err = r.channel.QueueBind(
		q.Name,     // queue name
		r.Key,      // 路由键
		r.Exchange, // exchange
		false,      //指示是否等待服务器的响应。在这里，设置为 false 表示要等待服务器的响应
		nil,
	)
	r.failOnErr("绑定失败", err)
	//获取接受消息的Delivery通道
	msgs, err := r.channel.Consume(
		q.Name, //表示要消费的队列名称
		"",     //表示消费者的名称，这里为空字符串。
		true,   //表示开启自动确认模式，即当消费者接收到消息后自动向 RabbitMQ 发送消息已被消费的回执(消息确认)
		false,  //表示不加锁。
		false,  //表示不禁止消费者使用同一连接发送消息。
		false,  //表示不等待接收返回的参数。
		nil,    //表示不传递其他参数。
	)
	r.failOnErr("创建消费者失败", err)
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func (r *Rabbit) ConsumeTopic() {
	err := r.channel.ExchangeDeclare(
		r.Exchange, //交换机名字
		"topic",    //交换机类型
		true,       //是否持久化
		false,      //是否自动删除，表示交换机在没有与之绑定的队列时是否自动删除
		false,      // 是否是内部交换机，表示该交换机是否被客户端使用
		false,      // 是否等待服务器响应，这里设置为 false。
		nil,
	)
	r.failOnErr("创建交换机失败", err)

	q, err := r.channel.QueueDeclare(
		"",    // name
		false, // 持久的
		false, //当队列不再使用时是否自动删除
		true,  //是否设置排他性队列（只能由声明它的连接使用）
		false, //是否设置为无等待（no-wait，等待服务器响应）。
		nil,
	)
	r.failOnErr("声明发送的消息队列失败", err)
	err = r.channel.QueueBind(
		q.Name,     // queue name
		r.Key,      // 路由键
		r.Exchange, // exchange
		false,      //指示是否等待服务器的响应。在这里，设置为 false 表示要等待服务器的响应
		nil,
	)
	r.failOnErr("绑定失败", err)
	//获取接受消息的Delivery通道
	msgs, err := r.channel.Consume(
		q.Name, //表示要消费的队列名称
		"",     //表示消费者的名称，这里为空字符串。
		true,   //表示开启自动确认模式，即当消费者接收到消息后自动向 RabbitMQ 发送消息已被消费的回执(消息确认)
		false,  //表示不加锁。
		false,  //表示不禁止消费者使用同一连接发送消息。
		false,  //表示不等待接收返回的参数。
		nil,    //表示不传递其他参数。
	)
	r.failOnErr("创建消费者失败", err)
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}

func (r *Rabbit) failOnErr(message string, err error) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}
