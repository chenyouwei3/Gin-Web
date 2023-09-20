package router

import (
	"github.com/gin-gonic/gin"
)

func GetEngine() *gin.Engine {
	engine := gin.Default()
	UserRouter(engine)
	RoleRouter(engine)
	ApiRouter(engine) //完
	return engine
}
