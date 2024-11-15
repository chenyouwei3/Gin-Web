package controller

import (
	"gin-web/utils/extendController"
	"github.com/gin-gonic/gin"
)

type DefaultController struct {
	extendController.BaseController
}

// HandleNotFound 404处理
func (d DefaultController) HandleNotFound(c *gin.Context) {
	d.SendNotFoundResponse(c)
}
