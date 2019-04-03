package models

import (
	"pkg/model"
	"github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	model.Model
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	JobId      int    `json:"job_id"`
	Content    string `json:"content"`
	ReplyFor   int    `json:"reply_for"` // 用户uid
}

func AddComment(comment Comment) error {
	if err := model.DB.Create(&comment).Error; err != nil {
		logrus.Errorf("models.AddComment error, err: %v", err.Error())
		return err
	}
	return nil
}

func DeleteCommentById(id int) error {
	if err := model.DB.Where("id = ?", id).Delete(Comment{}).Error; err != nil {
		logrus.Errorf("models.DeleteCommentById comment error, err: %v", err.Error())
		return err
	}
	return nil
}

func GetCommentsByJobId(jobId int) ([]*Comment, error) {
	var comments []*Comment
	err := model.DB.Order("created_on desc").Find(&comments).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("models.GetCommentsByJobId error, err: %v", err.Error())
		return nil, err
	}

	return comments, nil
}
