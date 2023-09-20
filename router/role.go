package router

import (
	"LoopyTicker/controller"
	"github.com/gin-gonic/gin"
)

func RoleRouter(engine *gin.Engine) {
	role := engine.Group("/role")
	{
		role.POST("/create", controller.CreateRole)
		role.DELETE("/deleted", controller.DeletedRole)
		role.PUT("/update", controller.UpdatedRole)
		role.GET("/get", controller.GetRole)
	}
}
