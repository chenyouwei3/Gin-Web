package controller

import (
	"errors"
	"gin-web/initialize/runLog"
	"gin-web/models/authcCenter"
	"gin-web/utils"
	"gin-web/utils/extendController"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RoleController struct {
	extendController.BaseController
}

type roleRequest struct {
	Role        authcCenter.Role
	AddApis     []int `json:"addApis"`
	DeletedApis []int `json:"deletedApis"`
}

func (r RoleController) Add(c *gin.Context) {
	var role roleRequest
	if err := c.Bind(&role); err != nil {
		runLog.ZapLog.Info("参数错误,role绑定错误" + err.Error())
		r.SendParameterErrorResponse(c, err)
		return
	}
	isExist, err := role.Role.IsExist()
	if isExist || err != nil {
		runLog.ZapLog.Info("数据重复")
		r.SendDataDuplicationResponse(c, err)
		return
	}
	if err = role.Role.Add(role.AddApis); err != nil {
		runLog.ZapLog.Info("添加role失败" + err.Error())
		r.SendCustomResponse(c, "添加role失败", "add role failed", err)
		return
	}
	r.SendSuccessResponse(c, "success")
}

func (r RoleController) Deleted(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		runLog.ZapLog.Info("参数错误,id为空")
		r.SendParameterErrorResponse(c, errors.New("参数错误,id为空"))
		return
	}
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		runLog.ZapLog.Info("id格式转化错误")
		r.SendParameterErrorResponse(c, errors.New("id转化错误为空"))
		return
	}
	//基于model
	if err := new(authcCenter.Role).Deleted(idInt64); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	r.SendSuccessResponse(c, "success")
}

func (r RoleController) Update(c *gin.Context) {
	var role roleRequest
	if err := c.Bind(&role); err != nil {
		runLog.ZapLog.Info("参数错误,role绑定错误" + err.Error())
		r.SendParameterErrorResponse(c, err)
		return
	}
	if err := role.Role.Update(role.AddApis, role.DeletedApis); err != nil {
		runLog.ZapLog.Info("更新role失败" + err.Error())
		r.SendCustomResponse(c, "更新role失败", "update role failed", err)
		return
	}
	r.SendSuccessResponse(c, "success")
}

func (r RoleController) GetAll(c *gin.Context) {
	var role authcCenter.Role
	role.Name = c.Query("name")
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		runLog.ZapLog.Info("参数错误,分页参数转化错误" + err.Error())
		r.SendParameterErrorResponse(c, err)
		return
	}
	resDB, err := role.GetAll(skip, limit, startTime, endTime)
	if err != nil {
		runLog.ZapLog.Info("查询role失败" + err.Error())
		r.SendCustomResponse(c, "查询role失败", "find role failed", err)
		return
	}
	r.SendSuccessResponse(c, resDB)
}
