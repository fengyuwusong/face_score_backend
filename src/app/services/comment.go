package services

import (
	"app/models"
	"github.com/sirupsen/logrus"
)

func AddComment(comment *models.Comment) error {
	err := models.AddComment(comment)
	if err != nil {
		logrus.Errorf("services.AddComment error, err: %v", err)
		return err
	}
	return nil
}

func DeleteCommentById(id int) error {
	err := models.DeleteCommentById(id)
	if err != nil {
		logrus.Errorf("services.DeleteCommentById error, err: %v", err)
		return err
	}
	return nil
}

func GetCommentByJobId(jobId int) ([]*models.Comment , error) {
	comments, err:= models.GetCommentsByJobId(jobId)
	if err != nil {
		logrus.Errorf("services.GetCommentByJobId error, err: %v", err)
		return nil, err
	}
	return comments, nil
}