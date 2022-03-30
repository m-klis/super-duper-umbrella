package producer

import (
	"os"

	"github.com/streadway/amqp"
)

type MessageBroker struct {
	Connect *amqp.Connection
	Channel *amqp.Channel
}

func InitializeMessageBroker() *MessageBroker {
	amqpHost := os.Getenv("AMQP_HOST")
	amqpUser := os.Getenv("AMQP_USER")
	amqpPass := os.Getenv("AMQP_PASS")
	amqpPort := os.Getenv("AMQP_PORT")
	link := "amqp://" + amqpUser + ":" + amqpPass + "@" + amqpHost + ":" + amqpPort + "/"

	ConnectRabbitMQ, err := amqp.Dial(link)
	if err != nil {
		panic(err)
	}

	ChannelRabbitMQ, err := ConnectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	var mb *MessageBroker = &MessageBroker{
		Connect: ConnectRabbitMQ,
		Channel: ChannelRabbitMQ,
	}
	return mb
}

func (mb *MessageBroker) CloseMessageBroker() {
	err := mb.Channel.Close()
	if err != nil {
		panic(err)
	}
	err = mb.Connect.Close()
	if err != nil {
		panic(err)
	}
}
