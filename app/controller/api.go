package controller

import (
	"github.com/gin-gonic/gin"
	"loopy-manager/app/model"
	"loopy-manager/app/service"
	"loopy-manager/pkg/e"
	"net/http"
)

type ApiController struct{}

func (A ApiController) CreateApi(c *gin.Context) {
	var api model.Api
	if err := c.Bind(&api); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.ApiService{}.CreateApi(api))
}

func (A ApiController) DeletedApi(c *gin.Context) {
	id := c.Query("id")
	if id == " " {
		c.JSON(e.ParameterError, e.GetMsg(e.ParameterError))
		return
	}
	c.JSON(http.StatusOK, service.ApiService{}.DeletedApi(id))
}

func (A ApiController) UpdatedApi(c *gin.Context) {
	var api model.Api
	if err := c.Bind(&api); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.ApiService{}.UpdateApi(api))
}

func (A ApiController) GetApi(c *gin.Context) {
	name := c.Query("name")
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	c.JSON(http.StatusOK, service.ApiService{}.GetApi(name, currPage, pageSize, startTime, endTime))
}
