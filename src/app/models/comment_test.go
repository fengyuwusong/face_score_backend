package models

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"pkg/config"
	"pkg/model"
)

func CommentTestSetUp() {
	config.InitConfig("../../../conf/face_score_backend.conf")
	model.Setup(config.GetConfig().Mysql)
	model.DB.Exec("TRUNCATE TABLE comment")
}

func CommentTestSetDown() {
	model.DB.Exec("TRUNCATE TABLE comment")
	model.CloseDB()
}

func TestAddComment(t *testing.T) {
	CommentTestSetUp()
	Convey("TestAddComment", t, func() {
		Convey("Success", func() {
			comment := Comment{
				UserId:  1,
				JobId:   1,
				Content: "TestAddComment",
			}
			err := AddComment(&comment)
			So(err, ShouldBeNil)
			So(comment.Id, ShouldEqual, 1)
		})
		Convey("Success with ReplyFor", func() {
			comment := Comment{
				UserId:   1,
				JobId:    1,
				Content:  "TestAddComment",
				ReplyFor: 1,
			}
			err := AddComment(&comment)
			So(err, ShouldBeNil)
			So(comment.Id, ShouldEqual, 2)
		})
	})
	CommentTestSetDown()
}


func TestDeleteCommentById(t *testing.T) {
	CommentTestSetUp()
	Convey("TestDeleteCommentById", t, func() {
		Convey("Success", func() {
			comment := Comment{
				UserId:   1,
				JobId:    1,
				Content:  "TestAddComment",
				ReplyFor: 1,
			}
			AddComment(&comment)
			commentId := 1
			err := DeleteCommentById(commentId)
			So(err, ShouldBeNil)
		})
		Convey("Not exist", func() {
			commentId := 1
			err := DeleteCommentById(commentId)
			So(err, ShouldNotBeNil)
		})

	})
	CommentTestSetDown()
}

func TestGetCommentsByJobId(t *testing.T) {
	CommentTestSetUp()
	comment := Comment{
		UserId:   1,
		JobId:    1,
		Content:  "TestAddComment",
		ReplyFor: 1,
	}
	AddComment(&comment)
	Convey("TestGetCommentsByJobId", t, func() {
		Convey("Success", func() {
			jobId := 1
			comments, err := GetCommentsByJobId(jobId)
			So(err, ShouldBeNil)
			So(len(comments), ShouldEqual, 1)
			So(comments[0].Id, ShouldEqual, 1)
		})
		Convey("Not exist", func() {
			jobId := 2
			comments, err := GetCommentsByJobId(jobId)
			So(err, ShouldBeNil)
			So(len(comments), ShouldEqual, 0)
		})
	})
	CommentTestSetDown()
}
