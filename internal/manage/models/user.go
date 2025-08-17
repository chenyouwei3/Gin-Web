package models

import (
	"errors"
	"fmt"
	mysqlDB "gin-web/init/mysql"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID         int64     `json:"id" gorm:"column:id;type:bigint;primaryKey;not null"`         // ID一般不从请求体绑定，不需要binding
	CreateTime time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime;index"`    // 自动创建，不需要binding
	UpdateTime time.Time `json:"updated_at" gorm:"column:updated_at;default:(-)"`             // 自动更新，不需要binding
	Name       string    `json:"name" gorm:"column:name;type:varchar(35);unique"`             // 名字必填，2~35长度
	Email      string    `json:"email" gorm:"column:email;type:varchar(35);not null"`         // 必填且为邮箱格式
	Account    string    `json:"account"  gorm:"column:account;type:varchar(35);not null"`    // 账号必填
	Password   string    `json:"password"  gorm:"column:password;type:varchar(100);not null"` // 密码必填且长度约束
	AvatarUrl  string    `json:"avatarUrl" gorm:"column:avatarUrl;type:varchar(50)"`          // 头像URL可选但格式必须是URL
	Roles      []Role    `json:"roles" gorm:"many2many:user_roles"`
}

func (User) tableName() string {
	return "users"
}

// 查询
func (u *User) GetList(skip, limit int, startTime, endTime string) ([]User, int64, error) {
	//总数
	var total int64
	countTx := mysqlDB.DB.Model(&User{})
	if startTime != "" && endTime != "" {
		countTx = countTx.Where("created_at >= ? AND created_at <= ?", startTime, endTime)
	}
	if u.Name != "" {
		countTx = countTx.Where("name LIKE ?", "%"+u.Name+"%") // 模糊查询 name
	}
	if u.Email != "" {
		countTx = countTx.Where("email LIKE ?", "%"+u.Email+"%") // 模糊查询 email
	}
	if err := countTx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	//子查询
	subQuery := mysqlDB.DB.Model(&User{}).Select("id").Order("created_at DESC")
	if startTime != "" && endTime != "" {
		subQuery = subQuery.Where("created_at >= ? AND created_at <= ?", startTime, endTime)
	}
	if u.Name != "" {
		subQuery = subQuery.Where("name LIKE ?", "%"+u.Name+"%") // 模糊查询 name
	}
	if u.Email != "" {
		subQuery = subQuery.Where("email LIKE ?", "%"+u.Email+"%") // 模糊查询 email
	}
	subQuery = subQuery.Offset(skip).Limit(limit)
	var resDB []User
	if err := mysqlDB.DB.Model(&User{}).
		Select("user.id", "name", "email", "account", "avatarUrl", "created_at", "updated_at").
		Joins("JOIN (?) AS tmp ON tmp.id = user.id", subQuery).
		Order("created_at DESC").
		Find(&resDB).Error; err != nil {
		return nil, 0, err
	}
	return resDB, total, nil
}

// 查询单一信息
func (u *User) GetOne() (*User, error) {
	var user User
	query := mysqlDB.DB.Model(&User{}).Preload("Roles", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name") // 只查 Role 的 id 和 name 字段
	})
	if u.Account != "" {
		query = query.Where("account = ?", u.Account)
	} else if u.Name != "" {
		query = query.Where("name = ?", u.Name)
	} else if u.Email != "" {
		query = query.Where("email = ?", u.Email)
	} else {
		return nil, errors.New("account 和 name 和 email 不能同时为空")
	}
	err := query.Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 插入
func (u *User) Insert(roleIds []int) error {
	return mysqlDB.DB.Transaction(func(tx *gorm.DB) error {
		// 创建用户
		if err := tx.Create(u).Error; err != nil {
			return err
		}
		if len(roleIds) == 0 {
			return nil // 事务成功，无需关联角色
		}
		// 查询所有要绑定的角色
		var roles []Role
		if err := tx.Find(&roles, roleIds).Error; err != nil {
			return err
		}
		fmt.Println(len(roles), len(roleIds))
		// 确保数量匹配，防止有无效 ID
		if len(roles) != len(roleIds) {
			return fmt.Errorf("角色数量不匹配")
		}

		// 添加用户与角色的关联关系（Many2Many）
		if err := tx.Model(u).Association("Roles").Append(roles); err != nil {
			return err
		}

		return nil
	})
}

// 修改
func (u *User) Update(addRoles, deletedRoles []int) error {
	u.UpdateTime = time.Now()
	err := mysqlDB.DB.Transaction(func(tx *gorm.DB) error {
		//更新用户基本信息
		if err := tx.Model(&User{}).Where("id = ?", u.ID).Updates(u).Error; err != nil {
			return fmt.Errorf("更新用户信息失败: %w", err)
		}
		// 删除关联
		if len(deletedRoles) > 0 {
			if err := tx.Table("user_roles").Where("user_id = ? AND role_id IN ?", u.ID, deletedRoles).Delete(nil).Error; err != nil {
				return fmt.Errorf("删除关联失败: %w", err)
			}
		}
		// 添加关联
		if len(addRoles) > 0 {
			records := make([]map[string]interface{}, len(addRoles))
			for i, roleId := range addRoles {
				records[i] = map[string]interface{}{
					"user_id": u.ID,
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

// 删除
func (u *User) Delete() error {
	return mysqlDB.DB.Transaction(func(tx *gorm.DB) error {
		// 清除 User 与 Roles 的关联关系
		err := tx.Model(&User{ID: u.ID}).Association("Roles").Clear()
		if err != nil {
			return err
		}
		// 删除 User 记录
		err = tx.Where("id = ?", u.ID).Delete(&User{}).Error
		if err != nil {
			return err
		}
		return nil
	})
}

// 是否存在
func (u *User) IsExist() (bool, error) {
	var exists bool
	err := mysqlDB.DB.Model(&Role{}).
		Select("1").
		Where("name = ?", u.Name).
		Limit(1).
		Scan(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

// 校验密码 输入的密码和传入的密码做校验
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// 设置密码
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}
