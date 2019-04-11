package job

import (
	"time"
	"app/models"
)

type JobInfo struct {
	models.Job
	Progress    int
	expiredTime int64
}

func NewJobInfo(job models.Job) *JobInfo {
	jobInfo := &JobInfo{
		Job:         job,
		Progress:    0,
		expiredTime: time.Now().Add(30 * time.Minute).UnixNano(),
	}
	return jobInfo
}

func (this *JobInfo) EndJobInfo(score int) {
	this.Job.Score = score
	this.Progress = 100
}

func (this *JobInfo) CanEliminate() bool {
	if this.expiredTime < time.Now().UnixNano() {
		return true
	}
	return false
}

func (j *JobInfo) Update() error {
	return nil
}

func (j *JobInfo) Reload() error {
	return nil
}

func (j *JobInfo) Close(bDrop bool) error {
	j.expiredTime = 0
	return nil
}
