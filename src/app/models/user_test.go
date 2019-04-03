package models

import (
	"pkg/config"
	"pkg/model"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func UserTestBegin() {
	config.InitConfig("F:\\code\\face_score_backend\\conf\\face_score_backend.conf")
	model.Setup(config.GetConfig().Mysql)
	model.DB.Delete(User{})
}

func UserTestFinish(){
	model.CloseDB()
}

func TestAddUser(t *testing.T) {
	Convey("TestAddUser", t, func() {

	})
}