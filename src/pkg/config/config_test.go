package config

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

// 测试配置读取
func TestConfig(t *testing.T) {
	Convey("TestConfig", t, func() {
		Convey("Success", func() {
			InitConfig("../../../conf/face_score_backend.conf")
			config = GetConfig()
			So(config.Mysql.Port, ShouldEqual, 3306)
		})

		Convey("Path Error", func(){
			err := InitConfig("...asd")
			So(err, ShouldNotBeNil)
		})
	})
}
