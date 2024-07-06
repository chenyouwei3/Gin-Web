package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	global2 "loopy-manager/app/global"
	"loopy-manager/app/model"
	"loopy-manager/pkg/auth"
	"loopy-manager/pkg/utils"
	"net/http"
	"strconv"
)

type ApiService struct{}

func (A ApiService) CreateApi(api model.Api) utils.Response {
	if err := global2.ApiTableMaster.Transaction(func(tx *gorm.DB) error {
		//查询角色重复
		var apiDB model.Api
		if err := tx.Select("url").Where("url = ?", api.Url).First(&apiDB).Error; (err != nil && !errors.Is(err, gorm.ErrRecordNotFound)) || api.Url == apiDB.Url {
			return fmt.Errorf("api重复:%w", err)
		}
		api.Id = global2.ApiSnowFlake.Generate().Int64()
		if err := tx.Create(&api).Error; err != nil {
			return fmt.Errorf("创建api失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("插入成功", "api")
}

func (A ApiService) DeletedApi(idString string) utils.Response {
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return utils.ErrorMess("字符串转化整数失败", err.Error())
	}
	if err := global2.ApiTableMaster.Transaction(func(tx *gorm.DB) error {
		tx0 := global2.RoleApiTableMaster.Begin()
		if err := tx0.Model(&model.Api{Id: id}).Association("Role").Clear(); err != nil {
			tx0.Rollback()
			return fmt.Errorf("清除关联失败:%w", err)
		}
		tx0.Commit()
		// 删除api记录
		if err := tx.Delete(&model.Api{}, id).Error; err != nil {
			return fmt.Errorf("删除api失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("删除事务失败", err.Error())
	}
	return utils.SuccessMess("删除成功", id)
}

func (A ApiService) UpdateApi(api model.Api) utils.Response {
	if err := global2.ApiTableMaster.Transaction(func(tx *gorm.DB) error {
		var apiDB model.Api
		if err := tx.Where("id = ?", api.Id).First(&apiDB).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("查询失败%w", err)
		}
		apiDB.Name = api.Name
		if err := tx.Save(&apiDB).Error; err != nil {
			return fmt.Errorf("更新api失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("修改api成功", api.Id)
}

func (A ApiService) GetApi(name, currPage, pageSize, startTime, endTime string) utils.Response {
	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		return utils.ErrorMess("数据转化失败", err.Error())
	}
	tx := global2.ApiTableSlave.ChooseSlave()
	if startTime != "" && endTime != "" {
		tx = tx.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var count int64
	var apiDB []model.Api
	res := tx.Debug().Where("name like ?", "%"+name+"%").Limit(limit).Offset(skip).Find(&apiDB).Count(&count)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}
	return utils.SuccessMess("成功", struct {
		Count int64       `json:"count" bson:"count"`
		Data  []model.Api `json:"data" bson:"data"`
	}{
		Count: count,
		Data:  apiDB,
	})
}

func LoginCookie(user model.User, c *gin.Context) utils.Response {
	var userDB model.User
	if err := global2.UserRoleTableSlave.ChooseSlave().Select("account").Where("account = ?", user.Account).First(&userDB).Error; err != nil {
		return utils.ErrorMess("账号已存在:", err)
	}
	//校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password+userDB.Salt)); err != nil {
		return utils.ErrorMess("密码错误:", err)
	}
	//查询角色信息
	var roleDB []model.Role
	if err := global2.RoleTableSlave.ChooseSlave().Select("id").Where("id = ?", user.RoleID).Find(&roleDB).Error; err != nil {
		return utils.ErrorMess("查询角色错误:", err)
	}
	//生成cookie
	cookie := http.Cookie{
		Name:   user.Account,                                //名称
		Value:  auth.CookieEncryption("cyw", user.Password), //值
		Path:   "/",                                         //有效路径
		Domain: c.ClientIP(),                                //cookie的有效域名
		//Expires:  time.Now().Add(time.Hour).UTC(), //过期时间
		MaxAge:   3600,
		HttpOnly: true, //js是否能够读取
	}
	//返回cookie
	//c.Writer.Header().Add("Set-Cookie", cookie.String())
	//设置在请求头上
	http.SetCookie(c.Writer, &cookie)
	res := map[string]interface{}{
		"id":       userDB.Id,
		"name":     userDB.Name,
		"account":  userDB.Account,
		"password": userDB.Password,
		"salt":     userDB.Salt,
		"role":     userDB.RoleID,
		//"cookie":   cookie,
	}
	return utils.SuccessMess("登陆成功", res)
}
