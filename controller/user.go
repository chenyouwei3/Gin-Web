package controller

import (
	"LoopyTicker/model"
	"LoopyTicker/service"
	"LoopyTicker/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// {
// "name":"陈幼伟",
// "account":"211010223",
// "password":"Cyw123456",
// "isValid":"true",
// "avatarUrl":"https://p3-pc.douyinpic.com/img/aweme-avatar/tos-cn-i-0813_e17affb4392f4fba9c458d3e2f14a4dc~c5_300x300.jpeg?from=2956013662",
// "sex":"男",
// "phone":"15329242550"
// }
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
