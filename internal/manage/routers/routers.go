package routers

import (
	"gin-web/init/runLog"
	"gin-web/internal/manage/controller"
	"gin-web/internal/manage/middleware"
	"gin-web/pkg/extendController"
	publicMiddleware "gin-web/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//store := cookie.NewStore([]byte("something-very-secret"))
	//r.Use(sessions.Sessions("mysession", store))
	r.Use(
		middleware.OperationLog(runLog.ZapLog),
		publicMiddleware.CorsMiddleware(),
		//publicMiddleware.AuthMiddleware(),
	)

	//r := gin.New()
	//r.Use(gin.Logger(), gin.Recovery()) //动记录所有 HTTP 请求的详细信息，如请求方法、请求路径、状态码、响应时间等。
	roleCH := &controller.RoleHandlerController{
		extendController.BaseController{
			RunLog: runLog.ZapLog,
		},
	}
	userCH := &controller.UserHandlerController{
		extendController.BaseController{
			RunLog: runLog.ZapLog,
		},
	}

	r.POST("/login", userCH.Login())
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, "success")
	})
	role := r.Group("role")
	{
		role.POST("/insert", roleCH.Insert()) //插入角色
		role.POST("/delete", roleCH.Delete()) //删除角色给
		role.POST("/update", roleCH.Update())
		role.GET("/getList", roleCH.GetList())
	}
	user := r.Group("user")
	{
		user.POST("/insert", userCH.Insert())
		user.POST("/delete", userCH.Delete())
		user.POST("/update", userCH.Update())
		user.GET("/getList", userCH.GetList())
		user.GET("/getUserByRoles", userCH.GetRolesByUserID())

	}
	return r
}
