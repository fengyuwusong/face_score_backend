package job

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"app/services"
	"strconv"
)

//查询任务接口
func Query(ctx *gin.Context) {
	jobIdStr := ctx.Param("jobId")
	if len(jobIdStr) == 0 {
		logrus.Errorf("query.Query error, jobId is empty")
		ctx.String(http.StatusBadRequest, "jobId is empty")
		ctx.Abort()
		return
	}
	jobId, err := strconv.Atoi(jobIdStr)
	if err != nil {
		logrus.Errorf("query.Query error, jobId not num")
		ctx.String(http.StatusBadRequest, "jobId not num")
		ctx.Abort()
		return
	}
	data, err := services.Query(jobId)
	if err != nil{
		logrus.Errorf("query.Query error, get job error")
		ctx.String(http.StatusInternalServerError, "get job error")
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, data)
}
