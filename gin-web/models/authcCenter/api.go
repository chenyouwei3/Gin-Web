package authcCenter

import (
	"fmt"
	mysqlDB "gin-web/initialize/mysql"
	"gin-web/utils"
	"strconv"
	"time"
)

type Api struct {
	Id         int64     `json:"id" gorm:"column:id;type:bigint;primaryKey;not null"`
	Name       string    `json:"name" gorm:"column:name;type:varchar(20);not null"` //Api名
	Url        string    ` json:"url" gorm:"column:url;type:varchar(20);not null"`  //地址
	Method     string    ` json:"method" gorm:"column:method;type:varchar(10);not null"`
	Desc       string    `json:"desc" gorm:"column:desc;type:varchar(10)"`           //描述
	CreateTime time.Time `json:"createTime" gorm:"column:createTime;autoCreateTime"` //创建time
	UpdateTime time.Time `json:"updateTime" gorm:"column:updateTime;autoCreateTime"` //修改time
}

func (a *Api) Add(api Api) error {
	api.CreateTime = time.Now()
	if err := mysqlDB.DB.Model(&Api{}).Create(api).Error; err != nil {
		return err
	}
	return nil
}

func (a *Api) Deleted(idString string) error {
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return err
	}
	if err := mysqlDB.DB.Delete(&Api{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (a *Api) Update(api Api) error {
	if api.Method != "POST" && api.Method != "GET" && api.Method != "DELETE" && api.Method != "PUT" {
		return fmt.Errorf("请求方法错误")
	}
	api.UpdateTime = time.Now()
	if err := mysqlDB.DB.Save(&api).Error; err != nil {
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
