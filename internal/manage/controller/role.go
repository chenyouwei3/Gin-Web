package controller

import (
	"gin-web/internal/manage/models"
	"gin-web/internal/manage/types"
	"gin-web/pkg"
	"gin-web/pkg/extendController"

	"github.com/gin-gonic/gin"
)

type RoleHandlerController struct {
	extendController.BaseController
}

// 查询角色列表
func (r *RoleHandlerController) GetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		//接收参数
		roleReq := &types.RoleGetListReq{
			Name:      c.Query("name"),
			CurrPage:  c.DefaultQuery("currPage", "1"),
			PageSize:  c.DefaultQuery("pageSize", "10"),
			StartTime: c.Query("startTime"),
			EndTime:   c.Query("endTime"),
		}
		skip, limit, err := pkg.GetPage(roleReq.CurrPage, roleReq.PageSize)
		if err != nil {
			r.SendServerErrorResponse(c, 5131, err)
			return
		}
		//DB操作
		roleDB := models.Role{
			Name: roleReq.Name,
		}
		resDB, total, err := roleDB.GetList(skip, limit, roleReq.StartTime, roleReq.EndTime)
		if err != nil {
			r.SendServerErrorResponse(c, 5130, err)
			return
		}
		r.SendSuccessResponse(c, types.RoleGetListResp{
			Roles: resDB,
			Total: total,
		})
	}
}

// 增加角色
func (r *RoleHandlerController) Insert() gin.HandlerFunc {
	return func(c *gin.Context) {
		var roleReq types.RoleInsertReq
		if err := c.Bind(&roleReq); err != nil {
			r.SendParameterErrorResponse(c, 4002, err)
			return
		}
		//DB操作
		roleDB := models.Role{
			Name: roleReq.Name,
			Desc: roleReq.Desc,
		}
		isExist, err := roleDB.IsExist()
		if isExist || err != nil {
			r.SendServerErrorResponse(c, 5101, err)
			return
		}
		if err = roleDB.Insert(); err != nil {
			r.SendServerErrorResponse(c, 5100, err)
			return
		}
		r.SendSuccessResponse(c, "success")
	}
}

// 删除角色
func (r *RoleHandlerController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		//参数校验
		var roleReq types.RoleDeleteReq
		err := c.Bind(&roleReq)
		if err != nil {
			r.SendParameterErrorResponse(c, 4002, err)
			return
		}
		//DB操作
		roleDB := models.Role{
			ID: roleReq.Id,
		}
		if err = roleDB.Delete(); err != nil {
			r.SendServerErrorResponse(c, 5110, err)
			return
		}
		r.SendSuccessResponse(c, "success")
	}
}

// 更新角色
func (r *RoleHandlerController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		//参数校验
		var roleReq types.RoleUpdateReq
		err := c.ShouldBind(&roleReq)
		if err != nil {
			r.SendParameterErrorResponse(c, 4002, err)
			return
		}
		//DB操作
		roleDB := models.Role{
			ID:   roleReq.Id,
			Name: roleReq.Name,
			Desc: roleReq.Desc,
		}
		if err = roleDB.Update(); err != nil {
			r.SendServerErrorResponse(c, 5120, err)
			return
		}
		r.SendSuccessResponse(c, "success")
	}
}
