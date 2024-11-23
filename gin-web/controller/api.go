package controller

import (
	"errors"
	"gin-web/initialize/runLog"
	"gin-web/models/authcCenter"
	"gin-web/utils"
	"gin-web/utils/extendController"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ApiController struct {
	extendController.BaseController
}

func (a ApiController) Add(c *gin.Context) {
	//接收参数并校验
	var api authcCenter.Api
	if err := c.Bind(&api); err != nil {
		runLog.ZapLog.Info("参数错误,api绑定错误" + err.Error())
		a.SendParameterErrorResponse(c, err)
		return
	}
	if api.Method != "POST" && api.Method != "GET" && api.Method != "DELETE" && api.Method != "PUT" {
		runLog.ZapLog.Info("参数错误,method方法错误")
		a.SendParameterErrorResponse(c, nil)
		return
	}
	isExist, err := api.IsExist()
	if isExist || err != nil {
		runLog.ZapLog.Info("数据重复")
		a.SendDataDuplicationResponse(c, err)
		return
	}
	if err = api.Add(); err != nil {
		runLog.ZapLog.Info("添加api失败" + err.Error())
		a.SendCustomResponse(c, "添加api失败", "add api failed", err)
		return
	}
	a.SendSuccessResponse(c, "success")
}

func (a ApiController) Deleted(c *gin.Context) {
	//接收参数并校验
	id := c.Query("id")
	if id == "" {
		runLog.ZapLog.Info("参数错误,id为空")
		a.SendParameterErrorResponse(c, errors.New("参数错误,id为空"))
		return
	}
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		runLog.ZapLog.Info("id格式转化错误")
		a.SendParameterErrorResponse(c, errors.New("id转化错误为空"))
		return
	}
	//基于model
	if err := new(authcCenter.Api).Deleted(idInt64); err != nil {
		runLog.ZapLog.Info("删除api失败")
		a.SendCustomResponse(c, "删除api失败", "deleted api failed", err)
		return
	}
	a.SendSuccessResponse(c, "success")
}

func (a ApiController) Update(c *gin.Context) {
	//接收参数并校验
	var api authcCenter.Api
	if err := c.Bind(&api); err != nil {
		runLog.ZapLog.Info("参数错误,api绑定错误" + err.Error())
		a.SendParameterErrorResponse(c, err)
		return
	}
	if api.Method != "POST" && api.Method != "GET" && api.Method != "DELETE" && api.Method != "PUT" {
		runLog.ZapLog.Info("参数错误,method方法错误")
		a.SendParameterErrorResponse(c, nil)
		return
	}
	if err := api.Update(); err != nil {
		runLog.ZapLog.Info("更新api失败" + err.Error())
		a.SendCustomResponse(c, "更新api失败", "update api failed", err)
		return
	}
	a.SendSuccessResponse(c, "success")
}

func (a ApiController) GetAll(c *gin.Context) {
	//接收参数并校验
	var api authcCenter.Api
	api.Name, api.Url = c.Query("name"), c.Query("url")
	currPage, pageSize := c.DefaultQuery("currPage", "1"), c.DefaultQuery("pageSize", "10")
	startTime, endTime := c.Query("startTime"), c.Query("endTime")
	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		runLog.ZapLog.Info("参数错误,分页参数转化错误" + err.Error())
		a.SendParameterErrorResponse(c, err)
		return
	}
	resDB, err := new(authcCenter.Api).GetAll(skip, limit, startTime, endTime)
	if err != nil {
		runLog.ZapLog.Info("查询api失败" + err.Error())
		a.SendCustomResponse(c, "查询api失败", "find api failed", err)
		return
	}
	a.SendSuccessResponse(c, resDB)
}
