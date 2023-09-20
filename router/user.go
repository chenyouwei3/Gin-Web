package router

import (
	"LoopyTicker/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(engine *gin.Engine) {
	user := engine.Group("user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
	}
}
