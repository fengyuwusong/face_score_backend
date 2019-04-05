package job

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"app/services"
)

// 提交任务接口
func Commit(ctx *gin.Context) {
	uidStr := ctx.PostForm("userId")
	if len(uidStr) == 0 {
		logrus.Errorf("commit.Commit error, error: uid is empty")
		ctx.String(http.StatusBadRequest, "uid is empty")
		ctx.Abort()
		return
	}
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		logrus.Errorf("commit.Commit error, error: uid is empty")
		ctx.String(http.StatusBadRequest, "uid is empty")
		ctx.Abort()
		return
	}

	fileIdStr := ctx.PostForm("fileId")
	if len(fileIdStr) == 0 {
		logrus.Errorf("commit.Commit error, error: fileId is empty")
		ctx.String(http.StatusBadRequest, "fileId is empty")
		ctx.Abort()
		return
	}
	fileId, err := strconv.Atoi(fileIdStr)
	if err != nil {
		logrus.Errorf("commit.Commit error, error: fileId is empty")
		ctx.String(http.StatusBadRequest, "fileId is empty")
		ctx.Abort()
		return
	}

	job, err := services.Commit(uid, fileId)
	if err != nil {
		logrus.Errorf("commit.Commit services.Commit error, error: %v", err)
		ctx.String(http.StatusBadRequest, "error: %v", err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, job)
}