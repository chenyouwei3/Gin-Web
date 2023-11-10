package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"loopy-manager/global"
	"loopy-manager/global/model"
	"loopy-manager/utils"
	"net/http"
)

func ApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取访问api的url和方法
		method, url := c.Request.Method, c.Request.URL.Path
		fmt.Println(url)
		var api model.Api

		res := global.ApiTable.Where("url = ? and method = ?", url, method).Find(&api)
		if res.Error != nil {
			fmt.Println(res.Error)
			c.Abort()
			return
		}

		//获取token解析出来的user
		userInterface, _ := c.Get("user")
		user := userInterface.(model.User)
		//获取user对应的role
		db := global.RoleTable
		_ = json.Unmarshal([]byte(user.RoleId), &user.RoleIds)
		var i int
		for i = 0; i < len(user.RoleIds); i++ {
			db = db.Or("id = ?", user.RoleIds[i])
		}
		var role []model.Role
		res = db.Limit(len(user.RoleIds)).Find(&role)
		if res.Error != nil {
			c.JSON(http.StatusOK, utils.ErrorMess("此用户无角色", nil))
			c.Abort()
			return
		}
		apiMap := make(map[int64]bool)
		var ok bool
		for i = 0; i < len(role); i++ {
			_ = json.Unmarshal([]byte(role[i].Api), &role[i].Apis)
			for j := range role[i].Apis {
				if _, ok = apiMap[role[i].Apis[j]]; !ok {
					apiMap[role[i].Apis[j]] = true
					_, ok = apiMap[api.Id] //判断权限是否存在
					if ok {
						c.Next()
						return
					}
				}
			}
		}

		c.JSON(http.StatusOK, utils.ErrorMess("验证api：此用户无访问此api的权限", nil))
		c.Abort()
		return
	}
}
