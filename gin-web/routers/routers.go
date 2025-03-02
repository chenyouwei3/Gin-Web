package routers

import (
	"fmt"
	"gin-web/controller"
	"gin-web/initialize/config"
	"gin-web/middleware"

	"github.com/gin-gonic/gin"
)

func RouterServerRun() {
	//注释是没有注释代码的参数
	//http.Handle()     //http.Handler()
	//http.HandleFunc() //	http.HandlerFunc()
	//r:=gin.Default()//自带gin.Logger()和gin.Recovery()两个中间件
	r := gin.New()
	r.Use(gin.Logger())   //动记录所有 HTTP 请求的详细信息，如请求方法、请求路径、状态码、响应时间等。
	r.Use(gin.Recovery()) //启用 错误恢复中间件，它会在出现 panic 错误时自动恢复，防止应用程序崩溃，并返回 HTTP 500 错误响应。

	//全局中间件注册
	r.Use(
		middleware.LimiterTokenBucket(5, 10),
		middleware.OperationLog("/Gin/V1.0", nil),
	)

	r.POST("/login")

	user := r.Group("/Gin/V1.0/user")
	{
		user.POST("/add", controller.UserController{}.Add)           //增加user
		user.DELETE("/deleted", controller.UserController{}.Deleted) //删除user
		user.PUT("/update", controller.UserController{}.Update)      //更新user
		user.GET("/getAll", controller.UserController{}.GetAll)      //查询user
	}
	role := r.Group("/Gin/V1.0/role")
	{
		role.POST("/add", controller.RoleController{}.Add)           //增加role
		role.DELETE("/deleted", controller.RoleController{}.Deleted) //删除role
		role.PUT("/update", controller.RoleController{}.Update)      //更新role
		role.GET("/getAll", controller.RoleController{}.GetAll)      //查询role
	}
	api := r.Group("/Gin/V1.0/api")
	{
		api.POST("/add", controller.ApiController{}.Add)           //增加api
		api.DELETE("/deleted", controller.ApiController{}.Deleted) //删除api
		api.PUT("/update", controller.ApiController{}.Update)      //更新api
		api.GET("/getAll", controller.ApiController{}.GetAll)      //查询api
	}
	// 捕捉不允许的方法
	r.NoMethod(controller.DefaultController{}.HandleNotFound) //无法匹配的方法
	r.NoRoute(controller.DefaultController{}.HandleNotFound)  //无法匹配的路由
	if err := r.Run(fmt.Sprintf("0.0.0.0:%d", config.Conf.APP.Port)); err != nil {
		panic(err)
	}
}
