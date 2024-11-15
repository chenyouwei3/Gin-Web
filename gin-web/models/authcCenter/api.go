package authcCenter

import (
	"errors"
	mysqlDB "gin-web/initialize/mysql"
	"gorm.io/gorm"
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

// 添加Api
func (a *Api) Add() error {
	a.CreateTime = time.Now()
	if err := mysqlDB.DB.Create(a).Error; err != nil {
		return err
	}
	return nil
}

// 删除Api
func (a *Api) Deleted(id int64) error {
	//启动事务
	return mysqlDB.DB.Transaction(func(tx *gorm.DB) error {
		// 清除 Api 与 Roles 的关联关系
		err := tx.Model(&Api{Id: id}).Association("Roles").Clear()
		if err != nil {
			return err
		}
		// 删除 Api 记录
		err = tx.Where("id = ?", id).Delete(&Api{}).Error
		if err != nil {
			return err
		}
		return nil
	})
}

// 修改Api
func (a *Api) Update() error {
	//修改参数
	a.UpdateTime = time.Now()
	if err := mysqlDB.DB.Updates(a).Error; err != nil {
		return err
	}
	return nil
}

func (a *Api) GetAll(name string, skip, limit int, startTime, endTime string) ([]Api, error) {
	tx := mysqlDB.DB
	if startTime != "" && endTime != "" {
		tx = tx.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var resDB []Api
	//Where("name like ?", "%"+name+"%")
	res := tx.Limit(limit).Offset(skip).Find(&resDB)
	if res.Error != nil {
		return nil, res.Error
	}
	return resDB, nil
}

// 查看是否存在
func (a *Api) IsExist() (bool, error) {
	//查重
	var api Api
	err := mysqlDB.DB.Model(&Api{}).Where("name = ? AND url = ?", a.Name, a.Url).Take(&api).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err // 其他错误
	}
	return true, nil // 存在
}
