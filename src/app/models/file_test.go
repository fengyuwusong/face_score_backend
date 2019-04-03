package models

import (
	"pkg/model"
	"pkg/config"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func FileTestSetUp() {
	config.InitConfig("../../../conf/face_score_backend.conf")
	model.Setup(config.GetConfig().Mysql)
	model.DB.Exec("TRUNCATE TABLE file")
}

func FileTestSetDown() {
	model.DB.Exec("TRUNCATE TABLE file")
	model.CloseDB()
}

func TestAddFile(t *testing.T) {
	FileTestSetUp()
	Convey("TestAddFile", t, func() {
		Convey("Success", func() {
			file := File{
				UserId: 1,
				Name:   "Test",
				Md5:    "md5",
				Uri:    "uri",
			}
			err := AddFile(&file)
			So(err, ShouldBeNil)
			So(file.Id, ShouldEqual, 1)
		})
		Convey("Exist error", func() {
			file := File{
				Id:     1,
				UserId: 1,
				Name:   "Test",
				Md5:    "md5",
				Uri:    "uri",
			}
			err := AddFile(&file)
			So(err, ShouldNotBeNil)
		})
	})
}

func TestGetFileById(t *testing.T) {
	FileTestSetUp()
	Convey("TestAddFile", t, func() {
		Convey("Success", func() {
			file := File{
				UserId: 1,
				Name:   "Test",
				Md5:    "md5",
				Uri:    "uri",
			}
			AddFile(&file)
			fileRes, err := GetFileById(1)
			So(err, ShouldBeNil)
			So(fileRes.Name, ShouldEqual, "Test")
		})
		Convey("Not exist error", func() {
			fileRes, err := GetFileById(2)
			So(err, ShouldNotBeNil)
			So(fileRes, ShouldBeNil)
		})
	})
	FileTestSetDown()
}

func TestGetFilesByUserId(t *testing.T) {
	FileTestSetUp()
	Convey("TestGetFilesByUserId", t, func() {
		Convey("Success", func() {
			file1 := File{
				UserId: 1,
				Name:   "Test1",
				Md5:    "md5",
				Uri:    "uri",
			}
			AddFile(&file1)
			file2 := File{
				UserId: 1,
				Name:   "Test2",
				Md5:    "md5",
				Uri:    "uri",
			}
			AddFile(&file2)
			files, err := GetFilesByUserId(1)
			So(err, ShouldBeNil)
			So(len(files), ShouldEqual, 2)
			So(files[0].Name, ShouldEqual, "Test1")
		})
		Convey("Not exist", func() {
			files, err := GetFilesByUserId(2)
			So(err, ShouldBeNil)
			So(len(files), ShouldEqual, 0)
		})
	})
	FileTestSetDown()
}
