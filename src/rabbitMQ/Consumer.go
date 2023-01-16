package main

import (
	"fmt"
	"log"

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

	// //生成队列
	// queue, err := channel.QueueDeclare("队列名", false, false, false, false, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	msgs, err := channel.Consume("队列名", "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	for v := range msgs {
		fmt.Println(string(v.Body))
	}
}
