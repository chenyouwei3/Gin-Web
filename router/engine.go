package router

import (
	"github.com/gin-gonic/gin"
	"loopy-manager/controller"
	"loopy-manager/middleware"
)

func GetEngine() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.Limiter(1, 1), middleware.Log(), middleware.Cors())
	engine.POST("/login", controller.Login)
	//engine.Use(middleware.JWTAuth(), middleware.ApiAuth())
	UserRouter(engine)
	RoleRouter(engine)
	ApiRouter(engine)
	return engine
}
