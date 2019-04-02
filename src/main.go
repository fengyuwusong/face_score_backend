package main

import (
	"pkg/config"
	"app/routers"
)

func main() {
	// 加载配置
	config.InitConfig("../conf/face_score_backend.conf")
	// 启动gin
	routers.Start()
}
