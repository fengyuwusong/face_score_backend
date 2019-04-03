package models

import (
	"pkg/model"
	"github.com/sirupsen/logrus"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Job struct {
	model.Model
	Id         int
	UserId     int
	FileId     int
	Score      int
	FinishedOn int
	Visible    bool
}

func AddJob(job *Job) error {
	if err := model.DB.Create(&job).Error; err != nil {
		logrus.Errorf("model.AddJob error, err: %v", err.Error())
		return err
	}
	return nil
}

// 任务完成
func EndJob(jobId int) error {
	e := model.DB.Table("job").Where("id = ?", jobId).Update("finished_on", time.Now().Unix())
	if err := e.Error; err != nil {
		logrus.Errorf("models.EndJob error, err: %v", err.Error())
		return err
	}
	if e.RowsAffected != 1{
		return errors.New("job id not exist")
	}
	return nil
}

// 可见
func Visible(jobId int) error {
	e := model.DB.Table("job").Where("id = ?", jobId).Update("visible", true)
	if err := e.Error; err != nil {
		logrus.Errorf("models.Visible error, err: %v", err.Error())
		return err
	}
	if e.RowsAffected != 1{
		return errors.New("job id not exist")
	}
	return nil
}

func GetJobByUserId(uid int) ([]*Job, error) {
	var jobs []*Job
	err := model.DB.Where("user_id = ?", uid).Find(&jobs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("model.GetJobByUserId error, err: %v", err.Error())
		return nil, err
	}
	return jobs, nil
}

// 获取所有可见的前10名
func GetJobsByRank10() ([]*Job, error) {
	var jobs []*Job
	err := model.DB.Where("visible = ?", true).Order("score desc").Limit(10).Find(&jobs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("model.GetJobsByRank10 error, err: %v", err.Error())
		return nil, err
	}
	return jobs, nil
}

// 随机获取可见的
func GetJobsByRandom(num int) ([]*Job, error) {
	var jobs []*Job
	err := model.DB.Where("visible = ?", true).Order("rand()").Limit(num).Find(&jobs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("model.GetJobsByRandom10 error, err: %v", err.Error())
		return nil, err
	}
	return jobs, nil
}

func GetJobById(jobId int) (*Job, error) {
	var job Job
	if err := model.DB.Where("id = ?", jobId).Find(&job).Error; err != nil {
		logrus.Errorf("models.GetJobById error, err: %v", err.Error())
		return nil, err
	}
	return &job, nil
}