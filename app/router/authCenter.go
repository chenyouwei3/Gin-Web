package router

import (
	"github.com/gin-gonic/gin"
	"loopy-manager/app/controller"
)

func AuthCenterRouter(engine *gin.Engine) {
	user := engine.Group("/user")
	{
		user.POST("/create", controller.UserController{}.CreateUser)     //增user*
		user.DELETE("/deleted", controller.UserController{}.DeletedUser) //删除user*
		user.PUT("/update", controller.UserController{}.UpdatedUser)     //修改*
		user.GET("/get", controller.UserController{}.GetUser)            //查找*
	}
	role := engine.Group("/role")
	{
		role.POST("/create", controller.RoleController{}.CreateRole)     //增role*
		role.DELETE("/deleted", controller.RoleController{}.DeletedRole) //删role*
		role.PUT("/update", controller.RoleController{}.UpdatedRole)     //改role******************
		role.GET("/get", controller.RoleController{}.GetRole)            //查role*
	}
	api := engine.Group("/api")
	{
		api.POST("/create", controller.ApiController{}.CreateApi)     //增api*
		api.DELETE("/deleted", controller.ApiController{}.DeletedApi) //删api*
		api.PUT("/update", controller.ApiController{}.UpdatedApi)     //改api*
		api.GET("/get", controller.ApiController{}.GetApi)            //查api*
	}
}
