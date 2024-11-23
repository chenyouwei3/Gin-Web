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

type userRequest struct {
	User         authcCenter.User `json:"user"`
	AddRoles     []int            `json:"addRoles"`
	DeletedRoles []int            `json:"deletedRoles"`
}

type UserController struct {
	extendController.BaseController
}

func (u UserController) Add(c *gin.Context) {
	var user userRequest
	if err := c.BindJSON(&user); err != nil {
		runLog.ZapLog.Info("参数错误,user绑定错误" + err.Error())
		u.SendParameterErrorResponse(c, err)
		return
	}
	isExist, err := user.User.IsExist()
	if isExist || err != nil {
		runLog.ZapLog.Info("数据重复")
		u.SendDataDuplicationResponse(c, err)
		return
	}
	if err = user.User.Add(user.AddRoles); err != nil {
		runLog.ZapLog.Info("添加user失败" + err.Error())
		u.SendCustomResponse(c, "添加user失败", "add user failed", err)
		return
	}
	u.SendSuccessResponse(c, "success")
}

func (u UserController) Deleted(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		runLog.ZapLog.Info("参数错误,id为空")
		u.SendParameterErrorResponse(c, errors.New("参数错误,id为空"))
		return
	}
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		runLog.ZapLog.Info("id格式转化错误")
		u.SendParameterErrorResponse(c, errors.New("id转化错误为空"))
		return
	}
	//基于model
	if err := new(authcCenter.User).Deleted(idInt64); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	u.SendSuccessResponse(c, "success")
}

func (u UserController) Update(c *gin.Context) {
	var user userRequest
	if err := c.Bind(&user); err != nil {
		runLog.ZapLog.Info("参数错误,user绑定错误" + err.Error())
		u.SendParameterErrorResponse(c, err)
		return
	}
	if err := user.User.Update(user.AddRoles, user.DeletedRoles); err != nil {
		runLog.ZapLog.Info("更新user失败" + err.Error())
		u.SendCustomResponse(c, "更新user失败", "update user failed", err)
		return
	}
	u.SendSuccessResponse(c, "success")
}

func (u UserController) GetAll(c *gin.Context) {
	var user authcCenter.User
	user.Name = c.Query("name")
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		runLog.ZapLog.Info("参数错误,分页参数转化错误" + err.Error())
		u.SendParameterErrorResponse(c, err)
		return
	}
	resDB, err := user.GetAll(skip, limit, startTime, endTime)
	if err != nil {
		runLog.ZapLog.Info("查询user失败" + err.Error())
		u.SendCustomResponse(c, "查询user失败", "find user failed", err)
		return
	}
	u.SendSuccessResponse(c, resDB)
}
