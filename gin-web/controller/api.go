package controller

import (
	"gin-web/models/authcCenter"
	"gin-web/utils"
	"gin-web/utils/extendController"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ApiController struct {
}

func (a ApiController) Add(c *gin.Context) {
	//接收参数并校验
	var api authcCenter.Api
	if err := c.Bind(&api); err != nil {
		c.JSON(10002, extendController.ErrBody)
		return
	}
	if api.Method != "POST" && api.Method != "GET" && api.Method != "DELETE" && api.Method != "PUT" {
		c.JSON(10002, extendController.ErrFormat)
		return
	}

	isExist, err := api.IsExist()
	if !isExist || err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if err = api.Add(); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "successful")
}

func (a ApiController) Deleted(c *gin.Context) {
	//接收参数并校验
	id := c.Query("id")
	if id == "" {
		c.JSON(10001, extendController.ErrQuery)
		return
	}
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(10003, extendController.ErrFormatConversion)
		return
	}
	//基于model
	if err := new(authcCenter.Api).Deleted(idInt64); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "successful")
}

func (a ApiController) Update(c *gin.Context) {
	//接收参数并校验
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

func (a ApiController) GetAll(c *gin.Context) {
	//接收参数并校验
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
	resDB, err := new(authcCenter.Api).GetAll(name, skip, limit, startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resDB)
}
