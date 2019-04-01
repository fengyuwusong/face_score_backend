package routers

import (
	"github.com/gin-gonic/gin"
	"app/config"
	"app/routers/apis/v1/user"
	"app/routers/apis/v1/file"
	"app/routers/apis/v1/job"
	"app/routers/apis/v1/comment"
	"app/middleware"
)

// 之后改进整合到app中
var engine *gin.Engine

func Start() {
	config := config.GetConfig()
	gin.SetMode(config.HttpServer.Mode)
	engine = gin.Default()

	// 注册中间件
	registerMiddleWare()

	// 注册路由
	registerRoutes()

}

func registerMiddleWare() {
	// 记录请求历史
	engine.Use(middleware.Entrance)
	// 账号认证
	engine.Use(middleware.Auth)
}

func registerRoutes() {
	// 查询用户信息
	engine.GET("/user", user.Get)
	// 添加用户
	engine.POST("/user", user.Add)
	// 上传文件
	engine.POST("/file", file.Upload)
	// 提交任务
	engine.POST("/commit/:method", job.Commit)
	// 查询任务
	engine.GET("/query/:jobid", job.Query)
	// 下载文件
	engine.GET("/file/:jobid/:fileid", file.Download)
	// 添加评论
	engine.POST("/comment/:userid", comment.Add)
	// 删除评论
	engine.DELETE("/comment/:commentid", comment.Delete)
	// 获取评论
	engine.GET("/comment/:userid", comment.Get)
}
