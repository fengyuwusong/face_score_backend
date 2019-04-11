package services

import (
	"app/models"
	"github.com/sirupsen/logrus"

	"app/job"
	"github.com/pkg/errors"
)

func Commit(userId, fileId int) (*models.Job, error) {
	jobModel := &models.Job{
		UserId: userId,
		FileId: fileId,
	}
	// 入库
	err := models.AddJob(jobModel)
	if err != nil {
		logrus.Errorf("services.Commit error, err: %v", err)
		return nil, err
	}
	// 写入缓存
	JPool := job.GetJPool()
	_, err = JPool.NewJobInfo(*jobModel)
	if err != nil {
		logrus.Errorf("services.Commit error, err: %v", err)
	}
	return jobModel, nil
}

func Query(jobId int) (*job.JobInfo, error) {
	JPool := job.GetJPool()
	jobInfo := JPool.GetJobInfo(jobId)
	if jobInfo == nil {
		logrus.Errorf("services.Query error, err: GetJobInfo return nil")
		return nil, errors.New("GetJobInfo return nil")
	}
	return jobInfo, nil
}

func GetJobsByUid(uid int) ([]*models.JobFull, error) {
	jobs, err := models.GetJobByUserId(uid)
	if err != nil {
		logrus.Errorf("services.GetJobsByUid error, err: %v", err)
		return nil, err
	}
	return jobs, nil
}

func SetVisible(jobId int) error {
	err := models.Visible(jobId)
	if err != nil {
		logrus.Errorf("services.SetVisible error, err: %v", err)
		return err
	}
	return nil
}

func GetJobsRank() ([]*models.JobFull, error) {
	jobs, err := models.GetJobsByRank(10)
	if err != nil {
		logrus.Errorf("services.GetJobsRank error, err: %v", err)
		return nil, err
	}
	return jobs, err
}

func GetJobsByRandom() ([]*models.JobFull, error) {
	jobs, err := models.GetJobsByRandom(10)
	if err != nil {
		logrus.Errorf("services.GetJobsByRandom error, err: %v", err)
		return nil, err
	}
	return jobs, err
}
