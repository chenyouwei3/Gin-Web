package models

import (
	mysqlDB "gin-web/init/mysql"
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID         int64     `json:"id"   gorm:"column:id;type:bigint;primaryKey;not null"` // ID 必须存在
	CreateTime time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime;index"`
	UpdateTime time.Time `json:"updated_at" gorm:"column:updated_at;default:(-)"`
	Name       string    `json:"name" gorm:"column:name;type:varchar(20);not null"`
	Desc       string    `json:"desc" gorm:"column:desc;type:varchar(20)" `
}

func (Role) tableName() string {
	return "role"
}

// 查询
func (r *Role) GetList(skip, limit int, startTime, endTime string) ([]Role, int64, error) {
	//总数
	var total int64
	countTx := mysqlDB.DB.Model(&Role{})
	if startTime != "" && endTime != "" {
		countTx = countTx.Where("created_at >= ? AND created_at <= ?", startTime, endTime)
	}
	if r.Name != "" {
		countTx = countTx.Where("name LIKE ?", "%"+r.Name+"%") // 模糊查询 name
	}
	if err := countTx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	//子查询
	subQuery := mysqlDB.DB.Model(&Role{}).Select("id").Order("created_at DESC")
	if startTime != "" && endTime != "" {
		subQuery = subQuery.Where("created_at >= ? AND created_at <= ?", startTime, endTime)
	}
	if r.Name != "" {
		subQuery = subQuery.Where("name LIKE ?", "%"+r.Name+"%") // 模糊查询 name
	}
	subQuery = subQuery.Offset(skip).Limit(limit)

	var resDB []Role

	if err := mysqlDB.DB.Model(&Role{}).
		//Select("id", "ip", "name", "url", "method", "desc", "created_at", "updated_at").
		//Preload("Apis").
		Joins("JOIN (?) AS tmp ON tmp.id = role.id", subQuery).
		Order("created_at DESC").
		Find(&resDB).Error; err != nil {
		return nil, 0, err
	}
	return resDB, total, nil
}

// 插入
func (r *Role) Insert() error {
	if err := mysqlDB.DB.Model(&Role{}).Create(r).Error; err != nil {
		return err
	}
	return nil
}

// 修改
func (r *Role) Update() error {
	r.UpdateTime = time.Now()
	if err := mysqlDB.DB.Model(&Role{}).
		Where("id = ?", r.ID).Updates(r).Error; err != nil {
		return err
	}
	return nil
}

// 删除
func (r *Role) Delete() error {
	//启动事务
	return mysqlDB.DB.Transaction(func(tx *gorm.DB) error {
		// 删除 Role 记录
		err := tx.Where("id = ?", r.ID).Delete(&Role{}).Error
		if err != nil {
			return err
		}
		return nil
	})
}

// 是否存在
func (r *Role) IsExist() (bool, error) {
	var exists bool
	err := mysqlDB.DB.Model(&Role{}).
		Select("1").
		Where("name = ?", r.Name).
		Limit(1).
		Scan(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}
