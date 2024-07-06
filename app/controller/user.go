package controller

import (
	"github.com/gin-gonic/gin"
	"loopy-manager/app/model"
	"loopy-manager/app/service"
	"loopy-manager/pkg/e"
	"net/http"
)

type UserController struct{}

func (U UserController) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.UserService{}.CreateUser(user))
}

func (U UserController) DeletedUser(c *gin.Context) {
	id := c.Query("id")
	if id == " " {
		c.JSON(e.ParameterError, e.GetMsg(e.ParameterError))
		return
	}
	c.JSON(http.StatusOK, service.UserService{}.DeletedUser(id))
}

func (U UserController) UpdatedUser(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.UserService{}.UpdatedUser(user))
}

func (U UserController) GetUser(c *gin.Context) {
	name := c.Query("name")
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	c.JSON(http.StatusOK, service.UserService{}.GetUser(name, currPage, pageSize, startTime, endTime))
}

func Login(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.LoginCookie(user, c))
}
