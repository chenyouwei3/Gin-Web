package authcCenter

import (
	"errors"
	"fmt"
	mysqlDB "gin-web/initialize/mysql"
	"gorm.io/gorm"
	"time"
)

type Role struct {
	Id         int64     `json:"id" gorm:"column:id;type:bigint;primaryKey;not null"`
	Name       string    `json:"name" gorm:"column:name;type:varchar(20);not null"`  //用户名
	Desc       string    `json:"desc" gorm:"column:desc;type:varchar(20)"`           //描述
	CreateTime time.Time `json:"createTime" gorm:"column:createTime;autoCreateTime"` //创建time
	UpdateTime time.Time `json:"updateTime" gorm:"column:updateTime;autoCreateTime"` //修改time
	User       []User    `gorm:"many2many:user_roles"`
	Apis       []Api     `gorm:"many2many:role_apis;"`
}

// 添加Role
func (r *Role) Add(apiIds []int) error {
	r.CreateTime = time.Now()
	return mysqlDB.DB.Transaction(func(tx *gorm.DB) error {
		// 创建 Role 记录
		res := tx.Create(r)
		if res.Error != nil {
			return res.Error
		}
		// 查找所有指定的 Api 记录
		var apis []Api
		if err := tx.Find(&apis, apiIds).Error; err != nil {
			return err
		}
		// 确保所有 apiIds 都存在
		if len(apis) != len(apiIds) {
			return fmt.Errorf("api数量不匹配")
		}
		// 关联 Api 到 Role
		if err := tx.Model(&Role{Id: r.Id}).Association("Apis").Append(apis); err != nil {
			return err
		}
		return nil
	})
}

// 删除Role
func (r *Role) Deleted(id int64) error {
	//删除role,受制于user/api
	return mysqlDB.DB.Transaction(func(tx *gorm.DB) error {
		// 清除 Api 与 Roles 的关联关系
		err := tx.Model(&Role{Id: id}).Association("Apis").Clear()
		if err != nil {
			return err
		}
		// 删除 Role 记录
		err = tx.Where("id = ?", id).Delete(&Role{}).Error
		if err != nil {
			return err
		}
		return nil
	})
}

// 修改Role
func (r *Role) Update(role Role) error {
	//修改参数
	role.UpdateTime = time.Now()
	if err := mysqlDB.DB.Save(&role).Error; err != nil {
		return err
	}
	return nil
}

func (r *Role) GetAll(name string, skip, limit int, startTime, endTime string) ([]Role, error) {
	tx := mysqlDB.DB
	if startTime != "" && endTime != "" {
		tx = tx.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var resDB []Role
	res := tx.Limit(limit).Offset(skip).Find(&resDB)
	if res.Error != nil {
		return nil, res.Error
	}
	return resDB, nil
}

// 查看是否存在
func (r *Role) IsExist() (bool, error) {
	//查重
	var role Role
	err := mysqlDB.DB.Model(&Role{}).Select("name").Where("name = ?", r.Name).Take(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil // 不存在
		}
		return false, err // 其他错误
	}
	return true, nil // 存在
}
