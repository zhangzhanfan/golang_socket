package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	//连接消息队列
	connection, err := amqp.Dial("amqp://admin:123@192.168.89.186:5672")
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	//创建通道
	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}
	//创建交换器
	// err = channel.ExchangeDeclare("el", "direct", true, false, false, true, nil)

	//生成队列
	queue, err := channel.QueueDeclare("队列名", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	//发送100条消息
	for i := 0; i < 100; i++ {
		err = channel.Publish("", queue.Name, false, false,
			amqp.Publishing{
				Timestamp:   time.Now(),
				ContentType: "text/plain",
				Body:        []byte("你好，消息队列，第一次用，哈哈哈哈哈哈"),
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}

}
