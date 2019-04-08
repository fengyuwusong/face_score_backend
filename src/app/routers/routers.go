package routers

import (
	"github.com/gin-gonic/gin"
	"app/routers/apis/v1/user"
	"app/routers/apis/v1/file"
	"app/routers/apis/v1/job"
	"app/routers/apis/v1/comment"
	"app/middleware"
	"pkg/httpservice"
	"pkg/config"
	"fmt"
)

func Start() {
	httpService := httpservice.HttpService{}

	// 启动
	httpService.Setup()

	// 注册中间件
	registerMiddleWare(httpService.Engine)

	// 注册路由
	registerRoutes(httpService.Engine)
	httpService.Engine.Run(fmt.Sprintf(":%d", config.GetConfig().HttpServer.Port))

}

func registerMiddleWare(engine *gin.Engine) {
	// 账号认证
	engine.Use(middleware.Auth)
}

func registerRoutes(engine *gin.Engine) {
	// 查询用户信息
	engine.GET("/user/:userId", user.Get)
	// 添加用户
	engine.POST("/user", user.Add)
	// 上传文件
	engine.POST("/file/:userId", file.Upload)
	// 添加评论
	engine.POST("/comment/:userId", comment.Add)
	// 删除评论
	engine.DELETE("/comment/:commentId", comment.Delete)
	// 获取评论
	engine.GET("/comment/:jobId", comment.Get)
	// 获取排行榜
	engine.GET("/job/rank", job.GetRank)
	// 随机获取job
	engine.GET("/job/random", job.GetByRandom)
	// 根据uid获取
	engine.GET("/job/uid/:userId", job.GetJobsByUid)
	// 提交任务
	engine.POST("/job/commit/:method", job.Commit)
	// 查询任务
	engine.GET("/job/query/:jobId", job.Query)
}
