package main

import "app/config"

func main() {
	// 加载配置
	config.InitConfig("../conf/face_score_backend.conf")
	// 启动gin

}
