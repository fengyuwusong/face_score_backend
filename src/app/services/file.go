package services

import (
	"mime/multipart"
	"os"
	"github.com/sirupsen/logrus"
	"app/models"
	"io"
	"pkg/utils"
	"strconv"
	"path/filepath"
)

func Upload(userId int, fileName string, fileBuffer multipart.File) (*models.File, error) {
	//rootPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	tempPath := "static/temp"
	// 判断文件夹是否存在
	if !utils.IsExist(tempPath){
		os.MkdirAll(tempPath, os.ModePerm)
	}
	//创建文件
	out, err := os.Create(tempPath + "/" + fileName)
	if err != nil {
		logrus.Errorf("services.Upload os.Create error, err: %v", err)
		return nil, err
	}
	defer out.Close()
	_, err = io.Copy(out, fileBuffer)
	if err != nil {
		logrus.Errorf("services.Upload io.Copy error, err: %v", err)
		return nil, err
	}
	// 计算文件md5
	md5 := utils.HashMD5(tempPath + "\\" + fileName)
	// 将文件转移到用户目录下 以md5命名
	userPath := "static/" + strconv.Itoa(userId)
	// 判断文件夹是否存在
	if !utils.IsExist(userPath){
		os.MkdirAll(userPath, os.ModePerm)
	}
	path := userPath + "/" + md5 + filepath.Ext(fileName)
	out, err = os.Create(path)
	if err != nil {
		logrus.Errorf("services.Upload os.Create error, err: %v", err)
		return nil, err
	}
	defer out.Close()
	// 将此处指针指回0, 重新读取进行复制
	fileBuffer.Seek(0, 0)
	_, err = io.Copy(out, fileBuffer)
	if err != nil {
		logrus.Errorf("services.Upload io.Copy  error, err: %v", err)
		return nil, err
	}
	// 将文件插入数据库
	file := &models.File{
		UserId: userId,
		Name:   fileName,
		Md5:    md5,
		Uri:    path,
	}
	err = models.AddFile(file)
	if err != nil {
		logrus.Errorf("services.Upload models.AddFile error, err: %v", err)
		return nil, err
	}
	return file, nil
}

