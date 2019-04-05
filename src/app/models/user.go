package models

import (
	"pkg/model"
	"github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

type User struct {
	model.Model
	Id       int
	OpenId   string
	UserName string
}

func AddUser(user *User) error {
	if err := model.DB.Create(&user).Error; err != nil {
		logrus.Errorf("models.AddUser error, err: %v", err.Error())
		return err
	}
	return nil
}

func GetUserById(id int) (*User, error) {
	var user User
	if err := model.DB.Where("id = ?", id).First(&user).Error; err != nil {
		// 用户不存在
		if err == gorm.ErrRecordNotFound {
			return nil, model.DataNotFound
		}
		logrus.Errorf("models.GetUserById error, err: %v", err.Error())
		return nil, err
	}
	return &user, nil
}

// check user
func CheckUserByOpenId(openId string) (bool, error) {
	var user User
	if err := model.DB.Where("open_id = ?", openId).Find(&user).Error; err != nil {
		// 用户不存在
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		logrus.Errorf("models.CheckUserByOpenId error, err: %v", err.Error())
		return false, err
	}
	return true, nil
}
