package pkg_amqp

import (
	"github.com/sirupsen/logrus"
)

func (this *MQBase) BindQueue(queueName string, callback ICallback) error {
	err := this.channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		logrus.Errorf("BindQueue error, err: %v", err)
		return err
	}
	msgs, err := this.channel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	// 启动协程监听消息
	go func() {
		for msg := range msgs {
			logrus.Infof("Received a message: %s", msg.Body)
			callback.Response(msg.MessageId, msg.Headers, msg.Body)
			msg.Ack(false)
		}
	}()
	return nil
}
