package controller

import (
	"github.com/gin-gonic/gin"
	"loopy-manager/global/model"
	"loopy-manager/service"
	"loopy-manager/utils"
	"net/http"
)

func Login(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorMess("登陆失败", err))
		return
	}
	c.JSON(http.StatusOK, service.Login(user))
}

func Register(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorMess("注册失败", err))
	}
	c.JSON(http.StatusOK, service.Register(user))
}
