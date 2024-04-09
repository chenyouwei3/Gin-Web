package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"loopy-manager/internal/model"
	"loopy-manager/internal/service"
	"loopy-manager/pkg/e"
	"net/http"
)

func AddComment(c *gin.Context) {
	var comment model.AddCommentForm
	if err := c.Bind(&comment); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	fmt.Println(comment, "controller")
	c.JSON(http.StatusOK, service.AddComment(comment))
}

func GetComment(c *gin.Context) {
	momentId := c.Query("mid")
	c.JSON(http.StatusOK, service.GetComment(momentId))
}

func AddMoment(c *gin.Context) {
	var moment model.AddMomentForm
	if err := c.Bind(&moment); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	fmt.Println(moment, "controller")
	c.JSON(http.StatusOK, service.AddMoment(moment))
}
