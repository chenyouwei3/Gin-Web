package controller

import (
	"gin-web/models/authcCenter"
	"gin-web/utils"
	"gin-web/utils/extendController"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RoleController struct {
}

type roleRequest struct {
	Role authcCenter.Role
	Apis []int `json:"apis"`
}

func (r RoleController) Add(c *gin.Context) {
	var role roleRequest
	if err := c.Bind(&role); err != nil {
		c.JSON(http.StatusInternalServerError, extendController.ErrBody)
		return
	}
	isExist, err := role.Role.IsExist()
	if isExist || err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if err = role.Role.Add(role.Apis); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "successful")
}

func (r RoleController) Deleted(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusInternalServerError, extendController.ErrQuery)
		return
	}
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, extendController.ErrFormatConversion)
		return
	}
	//基于model
	if err := new(authcCenter.Role).Deleted(idInt64); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "successful")
}

func (r *RoleController) Update(c *gin.Context) {
	var api authcCenter.Api
	if err := c.Bind(&api); err != nil {
		c.JSON(10002, extendController.ErrBody)
		return
	}
	if api.Method != "POST" && api.Method != "GET" && api.Method != "DELETE" && api.Method != "PUT" {
		c.JSON(10002, extendController.ErrFormat)
		return
	}

	if err := api.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "successful")
}

func (r *RoleController) GetAll(c *gin.Context) {
	name := c.Query("name")
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		c.JSON(10001, extendController.ErrQuery)
		return
	}

	resDB, err := new(authcCenter.Role).GetAll(name, skip, limit, startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resDB)
}
