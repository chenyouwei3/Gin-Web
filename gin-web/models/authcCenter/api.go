package authcCenter

import (
	"fmt"
	mysqlDB "gin-web/initialize/mysql"
	"gin-web/utils"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Api struct {
	Id         int64     `json:"id" gorm:"column:id;type:bigint;primaryKey;not null"`
	Name       string    `json:"name" gorm:"column:name;type:varchar(20);not null"` //Api名
	Url        string    ` json:"url" gorm:"column:url;type:varchar(20);not null"`  //地址
	Method     string    ` json:"method" gorm:"column:method;type:varchar(8);not null"`
	Desc       string    `json:"desc" gorm:"column:desc;type:varchar(20)"`           //描述
	CreateTime time.Time `json:"createTime" gorm:"column:createTime;autoCreateTime"` //创建time
	UpdateTime time.Time `json:"updateTime" gorm:"column:updateTime;autoCreateTime"` //修改time
	Roles      []Role    `gorm:"many2many:role_apis;"`                               //gorm结构体
}

func (a *Api) Add() error {
	a.CreateTime = time.Now()
	if err := mysqlDB.DB.Model(&Api{}).Create(a).Error; err != nil {
		return err
	}
	return nil
}

func (a *Api) Deleted(idString string) error {
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return err
	}
	//启动事务
	return mysqlDB.DB.Model(&Api{}).Transaction(func(tx *gorm.DB) error {
		// 查找 Api 记录并预加载关联的 Roles
		var api Api
		if err := tx.Preload("Roles").First(&api, id).Error; err != nil {
			return err
		}
		// 清除 Api 与 Roles 的关联关系
		if err := tx.Model(&api).Association("Roles").Clear(); err != nil {
			return err
		}
		// 删除 Api 记录
		if err := tx.Delete(&api).Error; err != nil {
			return err
		}
		return nil
	})
}

func (a *Api) Update(api Api) error {
	//修改参数
	if api.Method != "POST" && api.Method != "GET" && api.Method != "DELETE" && api.Method != "PUT" {
		return fmt.Errorf("请求方法错误")
	}
	api.UpdateTime = time.Now()

	if err := mysqlDB.DB.Model(&Api{}).Save(&api).Error; err != nil {
		return err
	}
	return nil
}

func (a *Api) GetAll(currPage, PageSize, startTime, endTime string) ([]Api, error) {
	skip, limit, err := utils.GetPage(currPage, PageSize)
	if err != nil {
		return nil, err
	}
	tx := mysqlDB.DB
	if startTime != "" && endTime != "" {
		tx = tx.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var count int64
	var resDB []Api
	res := tx.Model(&Api{}).Limit(limit).Offset(skip).Find(&resDB).Count(&count)
	if res.Error != nil {
		return nil, err
	}
	return resDB, nil
}

func (a *Api) IsExist() (bool, error) {
	//查重
	var count int64
	err := mysqlDB.DB.Model(&Api{}).Where("name = ? AND url = ?", a.Name, a.Url).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return true, nil
	}
	return false, fmt.Errorf("存在记录%d", count)
}
