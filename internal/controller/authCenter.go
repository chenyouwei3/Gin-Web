package controller

import (
	"github.com/gin-gonic/gin"
	"loopy-manager/internal/model"
	"loopy-manager/internal/service"
	"loopy-manager/pkg/e"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.CreateUser(user))
}

func DeletedUser(c *gin.Context) {
	id := c.Query("id")
	if id == " " {
		c.JSON(e.ParameterError, e.GetMsg(e.ParameterError))
		return
	}
	c.JSON(http.StatusOK, service.DeletedUser(id))
}

func UpdatedUser(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.UpdatedUser(user))
}

func GetUser(c *gin.Context) {
	name := c.Query("name")
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	c.JSON(http.StatusOK, service.GetUser(name, currPage, pageSize, startTime, endTime))
}

func CreateRole(c *gin.Context) {
	var role model.Role
	if err := c.Bind(&role); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.CreateRole(role))
}

func DeletedRole(c *gin.Context) {
	id := c.Query("id")
	if id == " " {
		c.JSON(e.ParameterError, e.GetMsg(e.ParameterError))
		return
	}
	c.JSON(http.StatusOK, service.DeletedRole(id))
}

func UpdatedRole(c *gin.Context) {
	var role model.Role
	if err := c.Bind(&role); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
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

func CreateApi(c *gin.Context) {
	var api model.Api
	if err := c.Bind(&api); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.CreateApi(api))
}

func DeletedApi(c *gin.Context) {
	id := c.Query("id")
	if id == " " {
		c.JSON(e.ParameterError, e.GetMsg(e.ParameterError))
		return
	}
	c.JSON(http.StatusOK, service.DeletedApi(id))
}

func UpdatedApi(c *gin.Context) {
	var api model.Api
	if err := c.Bind(&api); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
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

func Login(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.LoginCookie(user, c))
}
