package models

import (
	"pkg/model"
	"github.com/sirupsen/logrus"
)

type File struct {
	model.Model
	Id     int
	UserId int
	Name   string
	Md5    string
	Uri    string
}

func AddFile(file *File) error {
	if err := model.DB.Create(&file).Error; err != nil {
		logrus.Errorf("model.AddFile error, err: %v", err.Error())
		return err
	}
	return nil
}

func GetFileById(fileId int) (*File, error) {
	var file File
	if err := model.DB.Where("id = ?", fileId).Find(&file).Error; err != nil {
		logrus.Errorf("model.GetFile error, err: %v", err.Error())
		return nil, err
	}
	return &file, nil
}

func GetFilesByUserId(userId int) ([]*File, error) {
	var files []*File
	if err := model.DB.Where("user_id = ?", userId).Find(&files).Error; err != nil {
		logrus.Errorf("model.GetFile error, err: %v", err.Error())
		return nil, err
	}
	return files, nil
}
