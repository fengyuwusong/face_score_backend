package job

import (
	"github.com/streadway/amqp"
	"strconv"
	"pkg/pkg_amqp"
	"pkg/config"
	"fmt"
)
type Callback struct {
	MQConsumer pkg_amqp.MQBase
}

var CBack Callback
func SetUp()  {
	CBack := Callback{}
	mqConfig := config.GetConfig().RabbitMQ
	url := fmt.Sprintf("pkg_amqp://%s:%s@%s:%d",
		mqConfig.Username, mqConfig.Password, mqConfig.Host, mqConfig.Port)
	CBack.MQConsumer = pkg_amqp.SetUp(url)
	CBack.MQConsumer.BindQueue(config.GetConfig().RabbitMQ.ListenQueueName, CBack)
}

// 监听回调信息
func (this Callback) Response(msgId string, headers amqp.Table, body []byte) {
	jobId, _ := strconv.Atoi(msgId)
	score, _ := strconv.Atoi(string(body))
	GetJPool().EndJob(jobId, score)
}
