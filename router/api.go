package router

import (
	"github.com/gin-gonic/gin"
	"loopy-manager/controller"
)

func ApiRouter(engine *gin.Engine) {
	api := engine.Group("/api")
	{
		api.POST("/create", controller.CreateApi)
		api.DELETE("/deleted", controller.DeletedApi)
		api.PUT("/update", controller.UpdatedApi)
		api.GET("/get", controller.GetApi)
	}
}
