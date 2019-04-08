package pkg_amqp

import (
	"github.com/streadway/amqp"
	"github.com/sirupsen/logrus"
)

type MQBase struct {
	url        string
	connection *amqp.Connection // mq 连接
	channel    *amqp.Channel    // mq channel
}

func SetUp(url string) MQBase {
	conn, err := amqp.Dial(url)
	if err != nil {
		logrus.Errorf("Failed to connect to RabbitMQ, err: %v", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		logrus.Errorf("Failed to open a channel, err: %v", err)
	}
	mqBase := MQBase{
		url:        url,
		connection: conn,
		channel:    ch,
	}
	return mqBase
}

func (this *MQBase) Close() {
	if this.channel != nil {
		this.channel.Close()
		this.channel = nil
	}
	if this.connection != nil {
		this.connection.Close()
		this.connection = nil
	}
}
