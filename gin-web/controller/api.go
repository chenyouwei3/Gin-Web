package controller

import (
	"gin-web/models/authcCenter"
	"gin-web/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiController struct {
}

func (a *ApiController) Add(c *gin.Context) {
	var api authcCenter.Api
	if err := c.Bind(&api); err != nil {
		c.JSON(10002, response.ErrBody)
		return
	}
	if api.Method != "POST" && api.Method != "GET" && api.Method != "DELETE" && api.Method != "PUT" {
		c.JSON(10002, response.ErrBody)
		return
	}
	isExist, err := api.IsExist()
	if !isExist || err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err = api.Add()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "successful")
}
