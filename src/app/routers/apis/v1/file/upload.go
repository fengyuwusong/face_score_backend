package file

import (
	"github.com/gin-gonic/gin"
	"app/services"
	"net/http"
	"github.com/sirupsen/logrus"
	"strconv"
)

// 上传文件接口
func Upload(ctx *gin.Context) {
	uidStr := ctx.Param("userId")
	if len(uidStr) <= 0 {
		logrus.Errorf("file.Upload: not uid param.")
		ctx.String(http.StatusBadRequest, "not uid param.")
		ctx.Abort()
		return
	}

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		logrus.Errorf("file.Upload: uid not num.")
		ctx.String(http.StatusBadRequest, "uid not num.")
		ctx.Abort()
		return
	}

	fileBuffer, fileHeader, err := ctx.Request.FormFile("upload")
	if err != nil {
		logrus.Errorf("file.Upload: getFile error, err: %v", err)
		ctx.String(http.StatusBadRequest, "getFile error, err: %v", err)
		ctx.Abort()
		return
	}
	logrus.Infof("fileHeader: %v", fileHeader.Header)
	logrus.Infof("Filename: %v", fileHeader.Filename)
	file, err := services.Upload(uid, fileHeader.Filename, fileBuffer)
	if err != nil {
		logrus.Errorf("file.Upload: upload file error, err: %v", err)
		ctx.String(http.StatusBadRequest, "upload file error, err: %v", err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, *file)
}
