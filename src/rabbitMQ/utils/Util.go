package rabbitMQ

import (
	"log"

	"github.com/streadway/amqp"
)

func util() {
	//连接消息队列
	connection, err := amqp.Dial("amqp://admin:123@192.168.89.186:5672")
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

}
