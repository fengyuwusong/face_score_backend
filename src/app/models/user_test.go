package models

import (
	"pkg/config"
	"pkg/model"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func UserTestSetUp() {
	config.InitConfig("../../../conf/face_score_backend.conf")
	model.Setup(config.GetConfig().Mysql)
	model.DB.Exec("TRUNCATE TABLE user")
}

func UserTestSetDown() {
	model.DB.Exec("TRUNCATE TABLE user")
	model.CloseDB()
}

func TestAddUser(t *testing.T) {
	UserTestSetUp()
	Convey("TestAddUser", t, func() {
		// 成功案例
		Convey("Success", func() {
			user := User{
				OpenId:   "test",
				UserName: "test_name",
			}
			err := AddUser(&user)
			So(err, ShouldBeNil)
			So(user.Id, ShouldEqual, 1)
		})

		// 添加重复
		Convey("Add repeat error", func() {
			user := User{
				Id:       1,
				OpenId:   "test",
				UserName: "test_name",
			}
			err := AddUser(&user)
			So(err, ShouldNotBeNil)
		})
	})
	UserTestSetDown()
}

func TestGetUserById(t *testing.T) {
	UserTestSetUp()
	Convey("TestGetUserById", t, func() {
		Convey("Success", func() {
			user := User{
				Id:       1,
				OpenId:   "test",
				UserName: "test_name",
			}
			AddUser(&user)
			user1, err := GetUserById(1)
			So(err, ShouldBeNil)
			So(user1.Id, ShouldEqual, 1)
		})
		Convey("Not exist", func() {
			user1, err := GetUserById(2)
			So(err, ShouldNotBeNil)
			So(user1, ShouldBeNil)
		})
	})
}

func TestAuth(t *testing.T) {
	UserTestSetUp()
	Convey("TestAuth", t, func() {
		Convey("Success", func() {
			user := User{
				Id:       1,
				OpenId:   "test",
				UserName: "test_name",
			}
			AddUser(&user)
			exist, err := Auth("test")
			So(exist, ShouldBeTrue)
			So(err, ShouldBeNil)
		})
		Convey("Not exist", func() {
			exist, err := Auth("asfas")
			So(exist, ShouldBeFalse)
			So(err, ShouldBeNil)
		})
	})
	UserTestSetDown()
}
