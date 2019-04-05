package models

import (
	"pkg/config"
	"pkg/model"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func JobTestSetUp() {
	config.InitConfig("../../../conf/face_score_backend.conf")
	model.Setup(config.GetConfig().Mysql)
	model.DB.Exec("TRUNCATE TABLE job")
}

func JobTestSetDown() {
	model.DB.Exec("TRUNCATE TABLE job")
	model.CloseDB()
}

func TestAddJob(t *testing.T) {
	JobTestSetUp()
	Convey("TestAddJob", t, func() {
		Convey("Success", func() {
			job := Job{
				UserId: 1,
				FileId: 1,
				Score:  9,
			}
			err := AddJob(&job)
			So(err, ShouldBeNil)
			So(job.Id, ShouldEqual, 1)
		})
		Convey("With visible", func() {
			job := Job{
				UserId:  1,
				FileId:  1,
				Score:   9,
				Visible: true,
			}
			err := AddJob(&job)
			So(err, ShouldBeNil)
			So(job.Id, ShouldEqual, 2)
		})
		Convey("Exist", func() {
			job := Job{
				Id:      2,
				UserId:  1,
				FileId:  1,
				Score:   9,
				Visible: true,
			}
			err := AddJob(&job)
			So(err, ShouldNotBeNil)
		})
	})
	JobTestSetDown()
}

func TestEndJob(t *testing.T) {
	JobTestSetUp()
	Convey("TestEndJob", t, func() {
		Convey("Success", func() {
			job := Job{
				UserId: 1,
				FileId: 1,
			}
			AddJob(&job)
			err := EndJob(1, 10)
			jobRes, _ := GetJobById(1)
			So(err, ShouldBeNil)
			So(jobRes.FinishedOn, ShouldNotBeNil)
			So(jobRes.Score, ShouldEqual, 10)
		})
		Convey("Not exist", func() {
			err := EndJob(3, 4)
			So(err, ShouldNotBeNil)
		})
	})
	JobTestSetDown()
}

func TestVisible(t *testing.T) {
	JobTestSetUp()
	Convey("TestVisible", t, func() {
		Convey("Success", func() {
			job := Job{
				UserId: 1,
				FileId: 1,
				Score:  9,
			}
			AddJob(&job)
			err := Visible(1)
			So(err, ShouldBeNil)
			jobRes, _ := GetJobById(1)
			So(jobRes.Visible, ShouldBeTrue)
		})
		Convey("Not exist", func() {
			err := Visible(3)
			So(err, ShouldNotBeNil)
		})
	})
	JobTestSetDown()
}

func TestGetJobByUserId(t *testing.T) {
	JobTestSetUp()
	Convey("TestGetJobByUserId", t, func() {
		Convey("Success", func() {
			job := Job{
				UserId: 1,
				FileId: 1,
				Score:  9,
			}
			AddJob(&job)
			jobs, err := GetJobByUserId(1)
			So(err, ShouldBeNil)
			So(len(jobs), ShouldEqual, 1)
			So(jobs[0].Score, ShouldEqual, 9)
		})
		Convey("Not exist", func() {
			jobs, err := GetJobByUserId(3)
			So(err, ShouldBeNil)
			So(len(jobs), ShouldEqual, 0)
		})
	})
	JobTestSetDown()
}

func TestGetJobsByRank10(t *testing.T) {
	Convey("TestGetJobsByRank10", t, func() {
		JobTestSetUp()
		Convey("Success", func() {
			for i := 1; i <= 20; i++ {
				job := Job{
					UserId:  i,
					FileId:  i,
					Score:   i,
					Visible: true,
				}
				AddJob(&job)
			}
			jobs, err := GetJobsByRank(10)
			So(err, ShouldBeNil)
			So(len(jobs), ShouldEqual, 10)
			So(jobs[0].Score, ShouldEqual, 20)
			So(jobs[9].Score, ShouldEqual, 11)
		})
		Convey("Not enough 10", func() {
			for i := 1; i <= 5; i++ {
				job := Job{
					UserId:  i,
					FileId:  i,
					Score:   i,
					Visible: true,
				}
				AddJob(&job)
			}
			jobs, err := GetJobsByRank(10)
			So(err, ShouldBeNil)
			So(len(jobs), ShouldEqual, 5)
			So(jobs[0].Score, ShouldEqual, 5)
			So(jobs[4].Score, ShouldEqual, 1)
		})
		Convey("Not exist", func() {
			jobs, err := GetJobsByRank(10)
			So(err, ShouldBeNil)
			So(len(jobs), ShouldEqual, 0)
		})
		Reset(func() {
			JobTestSetDown()
		})
	})
}

func TestGetJobsByRandom10(t *testing.T) {
	JobTestSetUp()
	Convey("TestGetJobsByRandom10", t, func() {
		Convey("Success", func() {
			for i := 1; i <= 20; i++ {
				job := Job{
					UserId:  i,
					FileId:  i,
					Score:   i,
					Visible: true,
				}
				AddJob(&job)
			}
			jobs, err := GetJobsByRandom(10)
			So(err, ShouldBeNil)
			So(len(jobs), ShouldEqual, 10)
			sum1 := 0
			for _, val := range jobs {
				sum1 += val.Id
			}
			jobs, err = GetJobsByRandom(10)
			So(err, ShouldBeNil)
			So(len(jobs), ShouldEqual, 10)
			sum2 := 0
			for _, val := range jobs {
				sum2 += val.Id
			}
			So(len(jobs), ShouldEqual, 10)
			So(sum1, ShouldNotEqual, sum2)
		})
		Convey("Not enough 10", func() {
			jobs, err := GetJobsByRandom(22)
			So(err, ShouldBeNil)
			So(len(jobs), ShouldEqual, 20)
		})
	})
	JobTestSetDown()
}
