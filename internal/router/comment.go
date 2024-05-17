package router

import (
	"github.com/gin-gonic/gin"
	"loopy-manager/internal/controller"
)

func CommentRouter(engine *gin.Engine) {
	comment := engine.Group("/comments")
	{
		comment.POST("/create", controller.AddComment) //增comment
		comment.GET("/get", controller.GetComment)     //查找*
	}
	moment := engine.Group("/moment")
	{
		moment.POST("/create", controller.AddMoment)
	}
}
