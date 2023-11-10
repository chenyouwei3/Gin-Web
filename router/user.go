package router

import (
	"github.com/gin-gonic/gin"
	"loopy-manager/controller"
)

func UserRouter(engine *gin.Engine) {
	user := engine.Group("user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
	}
}
