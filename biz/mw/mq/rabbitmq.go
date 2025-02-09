package mq

import (
	"github.com/qingyggg/blog_server/pkg/constants"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ch *amqp.Channel
)

func Init() {
	var err error
	conn, err := amqp.Dial(constants.AmqpDSN)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	notifyInit()
}
