package controller

import (
	"github.com/gin-gonic/gin"
	"loopy-manager/app/model"
	"loopy-manager/app/service"
	"loopy-manager/pkg/e"
	"net/http"
)

type RoleController struct{}

func (R RoleController) CreateRole(c *gin.Context) {
	var role model.Role
	if err := c.Bind(&role); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.RoleService{}.CreateRole(role))
}

func (R RoleController) DeletedRole(c *gin.Context) {
	id := c.Query("id")
	if id == " " {
		c.JSON(e.ParameterError, e.GetMsg(e.ParameterError))
		return
	}
	c.JSON(http.StatusOK, service.RoleService{}.DeletedRole(id))
}

func (R RoleController) UpdatedRole(c *gin.Context) {
	var role model.Role
	if err := c.Bind(&role); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.RoleService{}.UpdateRole(role))
}

func (R RoleController) GetRole(c *gin.Context) {
	name := c.Query("name")
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	c.JSON(http.StatusOK, service.RoleService{}.GetRole(name, currPage, pageSize, startTime, endTime))
}
