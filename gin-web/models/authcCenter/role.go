package authcCenter

import (
	"fmt"
	mysqlDB "gin-web/initialize/mysql"
	"gin-web/utils"
	"strconv"
	"time"
)

type Role struct {
	Id         int64     `json:"id" gorm:"column:id;type:bigint;primaryKey;not null"`
	Name       string    `json:"name" gorm:"column:name;type:varchar(20);not null"`  //用户名
	Desc       string    `json:"desc" gorm:"column:desc;type:varchar(20)"`           //描述
	CreateTime time.Time `json:"createTime" gorm:"column:createTime;autoCreateTime"` //创建time
	UpdateTime time.Time `json:"updateTime" gorm:"column:updateTime;autoCreateTime"` //修改time
	User       []User    `gorm:"foreignKey:RoleID"`
	Apis       []Api     `gorm:"many2many:role_apis;"`
}

func (r *Role) Add(role Role) error {
	role.CreateTime = time.Now()
	if err := mysqlDB.DB.Model(&Role{}).Create(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *Role) Deleted(idString string) error {
	//删除role,受制于user/api
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return err
	}
	if err := mysqlDB.DB.Model(&Role{}).Delete(id).Error; err != nil {
		return err
	}
	return nil
}

func (r *Role) Update(role Role) error {
	//修改参数
	role.UpdateTime = time.Now()
	if err := mysqlDB.DB.Model(&Role{}).Save(&role).Error; err != nil {
		return err
	}
	return nil
}

func (r *Role) GetAll(currPage, PageSize, startTime, endTime string) ([]Role, error) {
	skip, limit, err := utils.GetPage(currPage, PageSize)
	if err != nil {
		return nil, err
	}
	tx := mysqlDB.DB
	if startTime != "" && endTime != "" {
		tx = tx.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var count int64
	var resDB []Role
	res := tx.Model(&Role{}).Limit(limit).Offset(skip).Find(&resDB).Count(&count)
	if res.Error != nil {
		return nil, err
	}
	return resDB, nil
}
func (r *Role) IsExist(name string) (bool, error) {
	//查重
	var count int64
	err := mysqlDB.DB.Model(&Role{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return true, nil
	}
	return false, fmt.Errorf("存在记录%d", count)
}
