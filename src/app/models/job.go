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

type JobFull struct {
	Job
	Uri string
}
func AddJob(job *Job) error {
	// todo 之后更新修改默认不可见 需掉用可见接口
	job.Visible = true
	if err := model.DB.Create(&job).Error; err != nil {
		logrus.Errorf("model.AddJob error, err: %v", err.Error())
		return err
	}
	return nil
}

// 任务完成
func EndJob(jobId int, score int) error {
	update := map[string]interface{}{
		"finished_on": time.Now().Unix(),
		"score": score,
	}
	e := model.DB.Table("job").Where("id = ?", jobId).Update(update)
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

func GetJobByUserId(uid int) ([]*JobFull, error) {
	var jobs []*JobFull
	err := model.DB.
		Table("job").
		Select("job.*, file.uri").
		Joins("left join file on file.id = job.file_id").
		Where("job.user_id = ?", uid).
		Find(&jobs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("model.GetJobByUserId error, err: %v", err.Error())
		return nil, err
	}
	return jobs, nil
}

// 获取所有可见的前10名
func GetJobsByRank(num int) ([]*JobFull, error) {
	var jobs []*JobFull
	err := model.DB.
		Table("job").
		Select("job.*, file.uri").
		Joins("left join file on file.id = job.file_id").
		Where("visible = ?", true).
		Order("score desc").
		Limit(num).
		Find(&jobs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("model.GetJobsByRank error, err: %v", err.Error())
		return nil, err
	}
	return jobs, nil
}

// 随机获取可见的
func GetJobsByRandom(num int) ([]*JobFull, error) {
	var jobs []*JobFull
	err := model.DB.
		Table("job").
		Select("job.*, file.uri").
		Joins("left join file on file.id = job.file_id").
		Where("visible = ?", true).
		Order("rand()").
		Limit(num).
		Find(&jobs).Error
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