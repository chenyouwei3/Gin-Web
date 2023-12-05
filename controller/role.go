package controller

import (
	"github.com/gin-gonic/gin"
	"loopy-manager/model"
	"loopy-manager/service"
	"net/http"
)

func CreateRole(c *gin.Context) {
	var role model.Role
	if err := c.Bind(&role); err != nil {
		c.JSON(http.StatusInternalServerError, "参数错误")
		return
	}
	c.JSON(http.StatusOK, service.CreateRole(role))
}

func DeletedRole(c *gin.Context) {
	id := c.Query("id")
	if id == " " {
		c.JSON(http.StatusInternalServerError, "参数错误")
		return
	}
	c.JSON(http.StatusOK, service.DeletedRole(id))
}

func UpdatedRole(c *gin.Context) {
	var role model.Role
	if err := c.Bind(&role); err != nil {
		c.JSON(http.StatusInternalServerError, "参数错误")
		return
	}
	c.JSON(http.StatusOK, service.UpdateRole(role))
}

func GetRole(c *gin.Context) {
	name := c.Query("name")
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	c.JSON(http.StatusOK, service.GetRole(name, currPage, pageSize, startTime, endTime))
}
