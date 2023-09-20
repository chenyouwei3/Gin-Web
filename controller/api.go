package controller

import (
	"LoopyTicker/model"
	"LoopyTicker/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateApi(c *gin.Context) {
	var api model.Api
	if err := c.Bind(&api); err != nil {
		c.JSON(http.StatusInternalServerError, "参数错误")
		return
	}
	c.JSON(http.StatusOK, service.CreateApi(api))
}

func DeletedApi(c *gin.Context) {
	id := c.Query("id")
	if id == " " {
		c.JSON(http.StatusInternalServerError, "参数错误")
		return
	}
	c.JSON(http.StatusOK, service.DeletedApi(id))
}

func UpdatedApi(c *gin.Context) {
	var api model.Api
	if err := c.Bind(&api); err != nil {
		c.JSON(http.StatusInternalServerError, "参数错误")
		return
	}
	c.JSON(http.StatusOK, service.UpdateApi(api))
}

func GetApi(c *gin.Context) {
	name := c.Query("name")
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	c.JSON(http.StatusOK, service.GetApi(name, currPage, pageSize, startTime, endTime))
}
