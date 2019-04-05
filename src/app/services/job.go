package services

import (
	"app/models"
	"github.com/sirupsen/logrus"
)

func Commit(userId, fileId int) (*models.Job, error) {
	job := &models.Job{
		UserId: userId,
		FileId: fileId,
	}
	// 获取file路径
	// todo 往mq发送消息

	// todo 将job写入内存缓存

	// 入库
	err := models.AddJob(job)
	if err != nil {
		logrus.Errorf("services.Commit error, err: %v", err)
		return nil, err
	}
	return job, nil
}

func Query(jobId int) (map[string]interface{}, error) {
	// todo 获取缓存判断是否已完成

	// todo 模拟进度条
	data := map[string]interface{}{"progress": 0}
	return data, nil
}

func GetJobsByUid(uid int) ([]*models.Job, error) {
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

func GetJobsRank() ([]*models.Job, error) {
	jobs, err := models.GetJobsByRank(10)
	if err != nil {
		logrus.Errorf("services.GetJobsRank error, err: %v", err)
		return nil, err
	}
	return jobs, err
}

func GetJobsByRandom() ([]*models.Job, error) {
	jobs, err := models.GetJobsByRandom(10)
	if err != nil {
		logrus.Errorf("services.GetJobsByRandom error, err: %v", err)
		return nil, err
	}
	return jobs, err
}