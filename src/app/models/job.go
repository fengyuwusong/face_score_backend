package models

import (
	"pkg/model"
	"github.com/sirupsen/logrus"
	"time"
	"github.com/jinzhu/gorm"
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

func AddJob(job Job) error {
	if err := model.DB.Create(&job).Error; err != nil {
		logrus.Errorf("model.AddJob error, err: %v", err.Error())
		return err
	}
	return nil
}

// 任务完成
func EndJob(jobId int) error {
	if err := model.DB.Table("job").Where("id = ?", jobId).Update("finished_on", time.Now().Unix()).Error; err != nil {
		logrus.Errorf("models.EndJob error, err: %v", err.Error())
		return err
	}
	return nil
}

// 可见
func Visible(jobId int) error {
	if err := model.DB.Table("job").Where("id = ?", jobId).Update("visible", true).Error; err != nil {
		logrus.Errorf("models.Visible error, err: %v", err.Error())
		return err
	}
	return nil
}

func GetJobByUid(uid int) ([]*Job, error) {
	var jobs []*Job
	err := model.DB.Where("user_id", uid).Find(jobs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("model.GetJobByUid error, err: %v", err.Error())
		return nil, err
	}
	return jobs, nil
}

// 获取所有可见的前10名
func GetJobsByRank10() ([]*Job, error) {
	var jobs []*Job
	err := model.DB.Where("visible", true).Order("score desc").Limit(10).Find(jobs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("model.GetJobsByRank10 error, err: %v", err.Error())
		return nil, err
	}
	return jobs, nil
}

// 随机获取可见的10名
func GetJobsByRandom10() ([]*Job, error) {
	var jobs []*Job
	err := model.DB.Where("visible", true).Order("rand()").Limit(10).Find(jobs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("model.GetJobsByRandom10 error, err: %v", err.Error())
		return nil, err
	}
	return jobs, nil
}
