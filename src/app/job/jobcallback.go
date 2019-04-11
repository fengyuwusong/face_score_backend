package job

import (
	"github.com/streadway/amqp"
	"strconv"
	"pkg/pkg_amqp"
	"pkg/config"
	"fmt"
	"github.com/sirupsen/logrus"
)
type Callback struct {
	MQConsumer pkg_amqp.MQBase
}

var CBack Callback
func SetUpJobCallback()  {
	CBack := Callback{}
	mqConfig := config.GetConfig().RabbitMQ
	url := fmt.Sprintf("amqp://%s:%s@%s:%d",
		mqConfig.Username, mqConfig.Password, mqConfig.Host, mqConfig.Port)
	CBack.MQConsumer = pkg_amqp.SetUp(url)
	CBack.MQConsumer.BindQueue(config.GetConfig().RabbitMQ.ListenQueueName, CBack)
}

// 监听回调信息
func (this Callback) Response(msgId string, headers amqp.Table, body []byte) {
	logrus.Info("get response message, msgId: " + msgId + "body: " + string(body))
	// todo msgid
	jobId, _ := strconv.Atoi(msgId)
	score, _ := strconv.Atoi(string(body))
	GetJPool().EndJob(jobId, score)
}
