package main

import (
	"pkg/config"
	"github.com/sirupsen/logrus"
	"app/job"
	"app/routers"
	"pkg/model"
)

func main() {
	// 加载配置
	err := config.InitConfig("conf/face_score_backend.conf")
	if err != nil {
		logrus.Fatal("InitConfig error, res: %v", err)
	}
	// 初始化models
	model.Setup(config.GetConfig().Mysql)
	// 初始化推送mq及缓存
	job.SetUpJobPool()
	// 绑定job消息队列
	job.SetUpJobCallback()
	// 启动gin
	routers.Start()
}
