package models

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"pkg/config"
	"pkg/model"
)

func CommentTestBegin() {
	config.InitConfig("F:\\code\\face_score_backend\\conf\\face_score_backend.conf")
	model.Setup(config.GetConfig().Mysql)
	model.DB.Delete(Comment{})
}

func CommentTestFinish(){
	model.CloseDB()
}

func TestAddComment(t *testing.T) {
	Convey("TestAddComment", t, func() {
		comment := Comment{
			UserId:  1,
			JobId:   1,
			Content: "TestAddComment",
		}
		err := AddComment(comment)
		So(err, ShouldBeNil)
	})
}

// 带ReplyFor
func TestAddCommentWithReplyFor(t *testing.T) {
	Convey("TestAddCommentWithReplyFor", t, func() {
		comment := Comment{
			UserId:   1,
			JobId:    1,
			Content:  "TestAddComment",
			ReplyFor: 1,
		}
		err := AddComment(comment)
		So(err, ShouldBeNil)
	})
}

func TestDeleteCommentById(t *testing.T) {
	Convey("TestDeleteCommentById", t, func() {
		commentId := 1
		err := DeleteCommentById(commentId)
		So(err, ShouldBeNil)
	})
}

func TestGetCommentsByJobId(t *testing.T) {
	Convey("TestGetCommentsByJobId", t, func() {
		jobId := 1
		comments, err := GetCommentsByJobId(jobId)
		So(err, ShouldBeNil)
		So(len(comments), ShouldEqual, 1)
		So(comments[0].Id, ShouldEqual, 2)
	})
}

func TestMain(m *testing.M) {
	CommentTestBegin()
	m.Run()
	CommentTestFinish()
}
