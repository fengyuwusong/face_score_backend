package main

import (
	"pkg/config"
	"app/routers"
	"github.com/sirupsen/logrus"
)

func main() {
	// 加载配置
	err := config.InitConfig("../conf/face_score_backend.conf")
	if err != nil {
		logrus.Fatal("InitConfig error, res: %v", err)
	}
	// 启动gin
	routers.Start()
}
