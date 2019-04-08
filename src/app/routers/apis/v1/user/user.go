package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"app/services"
	"app/models"
	"strconv"
)

func Get(ctx *gin.Context) {
	idStr := ctx.Param("userId")
	if len(idStr) == 0 {
		logrus.Errorf("user.Get error, error: id is empty")
		ctx.String(http.StatusBadRequest, "id is empty")
		ctx.Abort()
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil{
		logrus.Errorf("user.Get error, error: id is not num")
		ctx.String(http.StatusBadRequest, "id is not num")
		ctx.Abort()
		return
	}

	user, err := services.GetUserById(id)
	if err != nil{
		logrus.Errorf("user.Get error, error: %v", err)
		ctx.String(http.StatusInternalServerError, "err: %v", err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func Add(ctx *gin.Context) {
	// 获取参数并校验
	username := ctx.PostForm("username")
	if len(username) == 0 {
		logrus.Errorf("user.Add error, error: username is empty")
		ctx.String(http.StatusBadRequest, "username is empty")
		ctx.Abort()
		return
	}

	openId := ctx.PostForm("openId")
	if len(openId) == 0 {
		logrus.Errorf("user.Add error, error: openId is empty")
		ctx.String(http.StatusBadRequest, "openId is empty")
		ctx.Abort()
		return
	}
	user := &models.User{
		OpenId:   openId,
		UserName: username,
	}
	err := services.AddUser(user)
	if err != nil{
		logrus.Errorf("user.Add error, error: %v", err)
		ctx.String(http.StatusInternalServerError, "error: %v", err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, user)
}
