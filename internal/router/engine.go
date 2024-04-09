package router

import (
	"github.com/gin-gonic/gin"
	"loopy-manager/internal/controller"
	"loopy-manager/internal/middleware"
)

func GetEngine() *gin.Engine {
	engine := gin.Default()
	//路由日志*---*跨域
	engine.Use(middleware.OperationLogMiddleware(), middleware.CorsMiddleware())
	//限流
	engine.Use(middleware.LimiterBucket(1, 1))
	engine.POST("/login", controller.Login)
	//权限/jwt/cookie/session
	//engine.Use(middleware.AuthTokenMiddleware(), middleware.ApiAuth())
	//缓存
	//engine.Use(middleware.CacheTest())
	AuthCenterRouter(engine)
	CommentRouter(engine)
	return engine
}
