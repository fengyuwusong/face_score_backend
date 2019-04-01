package config

import "testing"

// 测试配置读取
func TestConfig(t *testing.T) {
	InitConfig("../../../conf/face_score_backend.conf")
	config = GetConfig()
	if config.HttpServer.Port != 80 {
		t.Error("config.HttpServerConfig.Port != 80")
	}
}
