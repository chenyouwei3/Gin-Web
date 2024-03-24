package middleware

import (
	"github.com/gin-gonic/gin"
	"loopy-manager/initialize/global"
	"loopy-manager/internal/model"
)

func ApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		method, url := c.Request.Method, c.Request.URL.Path
		var api model.Api
		if err := global.ApiTable.Select("url,method").Where("url= ? and method= ?", url, method).Take(&api).Error; err != nil {
		}
	}
}
