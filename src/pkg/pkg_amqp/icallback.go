package pkg_amqp

import "github.com/streadway/amqp"

// 消费者回调接口
type ICallback interface {
	Response(msgId string, headers amqp.Table, body []byte)
}
