package job

import (
	"github.com/gin-gonic/gin"
	"app/services"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetRank(ctx *gin.Context) {
	jobs, err := services.GetJobsRank()
	if err != nil {
		logrus.Errorf("job.GetRank error, err: %v", err)
		ctx.String(http.StatusBadRequest, "err: %v", err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, jobs)
}

func GetByRandom(ctx *gin.Context) {
	jobs, err := services.GetJobsByRandom()
	if err != nil {
		logrus.Errorf("job.GetByRandom error, err: %v", err)
		ctx.String(http.StatusBadRequest, "err: %v", err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, jobs)
}

func GetJobsByUid(ctx *gin.Context) {
	uidStr := ctx.Param("userId")
	if len(uidStr) <= 0 {
		logrus.Errorf("file.GetJobsByUid: not uid param.")
		ctx.String(http.StatusBadRequest, "not uid param.")
		ctx.Abort()
		return
	}

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		logrus.Errorf("file.GetJobsByUid: uid not num.")
		ctx.String(http.StatusBadRequest, "uid not num.")
		ctx.Abort()
		return
	}
	jobs, err := services.GetJobsByUid(uid)
	if err != nil {
		logrus.Errorf("job.GetJobsByUid error, err: %v", err)
		ctx.String(http.StatusBadRequest, "err: %v", err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, jobs)
}
