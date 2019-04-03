package models

import (
	"pkg/model"
	"pkg/config"
)

func FileTestInit() {
	config.InitConfig("F:\\code\\face_score_backend\\conf\\face_score_backend.conf")
	model.Setup(config.GetConfig().Mysql)
	model.DB.Delete(File{})
}

