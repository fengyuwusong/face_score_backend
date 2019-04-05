package services

import (
	"app/models"
	"github.com/sirupsen/logrus"
)

func AddUser(user *models.User) error {
	if err := models.AddUser(user); err != nil {
		logrus.Errorf("services.AddUser error, err: %v", err)
	}
	return nil
}

func GetUserById(id int) (*models.User, error) {
	user, err := models.GetUserById(id)
	if err != nil {
		logrus.Errorf("services.GetUserById error, err: %v", err)
		return nil, err
	}
	return user, nil
}

func CheckUserByOpenId(openId string) (bool, error) {
	exist, err := models.CheckUserByOpenId(openId)
	if err != nil{
		logrus.Errorf("services.CheckUserByOpenId error, err: %v", err)
		return false, err
	}
	return exist, nil
}