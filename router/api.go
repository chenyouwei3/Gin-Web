package router

import (
	"LoopyTicker/controller"
	"github.com/gin-gonic/gin"
)

func ApiRouter(engine *gin.Engine) {
	api := engine.Group("/api")
	{
		api.POST("/create", controller.CreateApi)
		api.DELETE("/deleted", controller.DeletedApi) //爆
		api.PUT("/update", controller.UpdatedApi)
		api.GET("/get", controller.GetApi)
	}
}
