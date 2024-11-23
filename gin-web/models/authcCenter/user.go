package authcCenter

import (
	"errors"
	"fmt"
	mysqlDB "gin-web/initialize/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id         int64     `json:"id" gorm:"column:id;type:bigint;primaryKey;not null"`
	Name       string    `json:"name" gorm:"column:name;type:varchar(20);not null"`         //用户名
	Account    string    `json:"account" gorm:"column:account;type:varchar(20);not null"`   //账号
	Password   string    `json:"password" gorm:"column:password;type:varchar(20);not null"` //密码
	AvatarUrl  string    `json:"avatarUrl" gorm:"column:avatarUrl;type:varchar(50)"`        //头像Url
	Sex        string    `json:"sex" gorm:"column:sex;type:varchar(3);not null"`            //性别
	Email      string    `json:"email" gorm:"column:email;type:varchar(20);not null"`       //邮箱
	Salt       string    `json:"salt" gorm:"column:salt;type:varchar(20);not null"`         //盐加密
	CreateTime time.Time `json:"createTime" gorm:"column:createTime;autoCreateTime"`        //创建time
	UpdateTime time.Time `json:"updateTime" gorm:"column:updateTime;default:(-)"`           //修改time
	Roles      []Role    `gorm:"many2many:user_roles"`                                      //外键role
}

func (u *User) Add(roleIds []int) error {
	u.CreateTime = time.Now()
	return mysqlDB.DB.Transaction(func(tx *gorm.DB) error {
		res := tx.Create(u)
		if res.Error != nil {
			return res.Error
		}
		// 查找所有指定的 roles 记录
		var roles []Role
		if err := tx.Find(&roles, roleIds).Error; err != nil {
			return err
		}
		// 确保所有 roles 都存在
		if len(roles) != len(roleIds) {
			return fmt.Errorf("role数量不匹配")
		}
		// 关联 Role 到 User
		if err := tx.Model(&User{Id: u.Id}).Association("Roles").Append(roles); err != nil {
			return err
		}
		return nil
	})
}

func (u *User) Deleted(id int64) error {
	//删除role,受制于user/api
	return mysqlDB.DB.Transaction(func(tx *gorm.DB) error {
		// 清除 User 与 Roles 的关联关系
		err := tx.Model(&User{Id: id}).Association("Roles").Clear()
		if err != nil {
			return err
		}
		// 删除 User 记录
		err = tx.Where("id = ?", id).Delete(&User{}).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (u *User) Update(addRoles, deletedRoles []int) error {
	u.UpdateTime = time.Now()
	err := mysqlDB.DB.Transaction(func(tx *gorm.DB) error {
		//更新用户基本信息
		if err := tx.Model(u).Save(u).Error; err != nil {
			return fmt.Errorf("更新角色信息失败: %w", err)
		}
		// 删除关联
		if len(deletedRoles) > 0 {
			if err := tx.Table("user_roles").Where("role_id = ? AND user_id IN ?", u.Id, deletedRoles).Delete(nil).Error; err != nil {
				return fmt.Errorf("删除关联失败: %w", err)
			}
		}
		// 添加关联
		if len(addRoles) > 0 {
			records := make([]map[string]interface{}, len(addRoles))
			for i, roleId := range addRoles {
				records[i] = map[string]interface{}{
					"user_id": u.Id,
					"role_id": roleId,
				}
			}
			if err := tx.Table("user_roles").Create(records).Error; err != nil {
				return fmt.Errorf("添加关联失败: %w", err)
			}
		}
		return nil
	})
	return err
}

func (u *User) GetAll(skip, limit int, startTime, endTime string) ([]User, error) {
	tx := mysqlDB.DB
	if startTime != "" && endTime != "" {
		tx = tx.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var resDB []User
	res := tx.Model(&User{}).Limit(limit).Offset(skip).Find(&resDB)
	if res.Error != nil {
		return nil, res.Error
	}
	return resDB, nil
}

func (u *User) IsExist() (bool, error) {
	//查重
	var user User
	err := mysqlDB.DB.Model(&User{}).Where("account = ? OR name= ?", u.Account, u.Name).Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil // 记录不存在
	}
	if err != nil {
		return false, err // 其他错误
	}
	return true, nil // 记录存在
}

func (u *User) ChangePassword() {

}
