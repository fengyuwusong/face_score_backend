package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"app/models"
	"app/services"
)

// 添加评论
func Add(ctx *gin.Context) {
	// 获取参数并校验
	uidStr := ctx.PostForm("userId")
	if len(uidStr) == 0 {
		logrus.Errorf("comment.Add error, error: userId is empty")
		ctx.String(http.StatusBadRequest, "userId is empty")
		ctx.Abort()
		return
	}
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		logrus.Errorf("comment.Add error, error: userId is not num")
		ctx.String(http.StatusBadRequest, "userId is not num")
		ctx.Abort()
		return
	}

	jobIdStr := ctx.PostForm("jobId")
	if len(jobIdStr) == 0 {
		logrus.Errorf("comment.Add error, error: jobId is empty")
		ctx.String(http.StatusBadRequest, "jobId is empty")
		ctx.Abort()
		return
	}
	jobId, err := strconv.Atoi(jobIdStr)
	if err != nil {
		logrus.Errorf("comment.Add error, error: jobId is not num")
		ctx.String(http.StatusBadRequest, "jobId is not num")
		ctx.Abort()
		return
	}

	content := ctx.PostForm("content")
	if len(content) == 0 {
		logrus.Errorf("comment.Add error, error: content is empty")
		ctx.String(http.StatusBadRequest, "content is empty")
		ctx.Abort()
		return
	}

	replyForStr := ctx.PostForm("replyFor")
	var comment *models.Comment

	if len(replyForStr) == 0 {
		comment = &models.Comment{
			UserId:  uid,
			JobId:   jobId,
			Content: content,
		}
	} else {
		replyFor, err := strconv.Atoi(replyForStr)
		if err != nil {
			logrus.Errorf("comment.Add error, error: replyFor is not num")
			ctx.String(http.StatusBadRequest, "replyFor is not num")
			ctx.Abort()
			return
		}
		comment = &models.Comment{
			UserId:   uid,
			JobId:    jobId,
			Content:  content,
			ReplyFor: replyFor,
		}
	}

	err = services.AddComment(comment)
	if err != nil {
		logrus.Errorf("comment.Add services.AddComment error, error: %v", err)
		ctx.String(http.StatusBadRequest, "error: %v", err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// 删除评论
func Delete(ctx *gin.Context) {
	idStr := ctx.Param("commentId")
	if len(idStr) == 0 {
		logrus.Errorf("comment.Delete error, error: id is empty")
		ctx.String(http.StatusBadRequest, "id is empty")
		ctx.Abort()
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.Errorf("comment.Delete error, error: id is not num")
		ctx.String(http.StatusBadRequest, "id is not num")
		ctx.Abort()
		return
	}

	err = services.DeleteCommentById(id)
	if err != nil {
		logrus.Errorf("comment.Delete services.DeleteCommentById error, error: %v", err)
		ctx.String(http.StatusBadRequest, "error: %v", err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, "")
}

// 获取评论
func Get(ctx *gin.Context) {
	jobIdStr := ctx.Param("jobId")
	if len(jobIdStr) == 0 {
		logrus.Errorf("comment.Get error, error: jobId is empty")
		ctx.String(http.StatusBadRequest, "jobId is empty")
		ctx.Abort()
		return
	}
	jobId, err := strconv.Atoi(jobIdStr)
	if err != nil {
		logrus.Errorf("comment.Get error, error: jobId is not num")
		ctx.String(http.StatusBadRequest, "jobId is not num")
		ctx.Abort()
		return
	}
	jobs, err := services.GetCommentByJobId(jobId)
	if err != nil {
		logrus.Errorf("comment.Get services.GetCommentByJobId error, error: %v", err)
		ctx.String(http.StatusBadRequest, "error: %v", err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, jobs)
}
