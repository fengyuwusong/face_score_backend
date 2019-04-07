package job

import (
	"pkg/cache"
	"fmt"
	"github.com/pkg/errors"
	"app/models"
	"pkg/pkg_amqp"
	"pkg/config"
	"github.com/sirupsen/logrus"
	"strconv"
	"math/rand"
)

type JobInfoPool struct {
	cache.CachePoolBase
	MQProducer pkg_amqp.MQBase
}

var JPool JobInfoPool

func init() {
	JPool = JobInfoPool{}
	mqConfig := config.GetConfig().RabbitMQ
	url := fmt.Sprintf("pkg_amqp://%s:%s@%s:%d",
		mqConfig.Username, mqConfig.Password, mqConfig.Host, mqConfig.Port)
	JPool.MQProducer = pkg_amqp.SetUp(url)
}

func GetJPool() JobInfoPool {
	return JPool
}

func (j JobInfoPool) NewJobInfo(job models.Job) (*JobInfo, error) {
	jobInfo := NewJobInfo(job)
	j.Add(fmt.Sprintf("%d_job", jobInfo.Id), jobInfo)
	// 查询fileid所在路径
	file, err := models.GetFileById(job.FileId)
	if err != nil {
		logrus.Errorf("NewJobInfo models.GetFileById error, err: %v", err)
		return nil, err
	}
	body := []byte(file.Uri)
	err = j.MQProducer.Push2MQ(strconv.Itoa(job.Id), body, config.GetConfig().RabbitMQ.PushQueueName)
	if err != nil {
		logrus.Errorf("NewJobInfo MQProducer.Push2MQ error, err: %v", err)
		return nil, err
	}
	return &jobInfo, nil
}

func (j JobInfoPool) GetJobInfo(jobId int) *JobInfo {
	if r, b := j.Get(fmt.Sprintf("%d_job", jobId)); b {
		jobInfo := r.(*JobInfo)
		if jobInfo.Progress != 100{
			jobInfo.Progress += rand.Intn(5)
			if jobInfo.Progress>= 100{
				jobInfo.Progress = 99
			}
		}
		return jobInfo
	}

	// 不用分布式
	//mJobInfo, err := JobDB.Query(jobId)
	//if err != nil {
	//	return nil
	//}
	//
	//if mJobInfo == nil {
	//	return nil
	//}

	return nil
}

func (j JobInfoPool) EndJob(jobId, score int) error {
	r, b := j.Get(fmt.Sprintf("%d_job", jobId))
	if !b {
		return errors.New("can't find jobId")
	}
	jobInfo := r.(JobInfo)
	jobInfo.EndJobInfo(score)
	return nil
}
