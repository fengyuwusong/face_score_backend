package httpservice

import (
	"pkg/config"
	"github.com/gin-gonic/gin"
)

type HttpService struct {
	Engine *gin.Engine
}

func (this *HttpService) Setup() {
	cfg := config.GetConfig()
	gin.SetMode(cfg.HttpServer.Mode)
	this.Engine = gin.Default()
}
