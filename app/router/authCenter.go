package router

import (
	"github.com/gin-gonic/gin"
	"loopy-manager/app/controller"
)

func AuthCenterRouter(engine *gin.Engine) {
	user := engine.Group("/user")
	{
		user.POST("/create", controller.CreateUser)     //增user*
		user.DELETE("/deleted", controller.DeletedUser) //删除user*
		user.PUT("/update", controller.UpdatedUser)     //修改*
		user.GET("/get", controller.GetUser)            //查找*
	}
	role := engine.Group("/role")
	{
		role.POST("/create", controller.CreateRole)     //增role*
		role.DELETE("/deleted", controller.DeletedRole) //删role*
		role.PUT("/update", controller.UpdatedRole)     //改role******************
		role.GET("/get", controller.GetRole)            //查role*
	}
	api := engine.Group("/api")
	{
		api.POST("/create", controller.CreateApi)     //增api*
		api.DELETE("/deleted", controller.DeletedApi) //删api*
		api.PUT("/update", controller.UpdatedApi)     //改api*
		api.GET("/get", controller.GetApi)            //查api*
	}
}
