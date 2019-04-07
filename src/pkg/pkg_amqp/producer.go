package pkg_amqp

import (
	"github.com/streadway/amqp"
	"github.com/sirupsen/logrus"
)

func (this *MQBase) Push2MQ(msgId string, body []byte, queueName string) error {
	err := this.channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/octet-stream",
			Body:         body,
			MessageId:    msgId,
		})
	if err != nil {
		logrus.Errorf("Push2MQ error, err: %v", err)
		return err
	}
	return nil
}
