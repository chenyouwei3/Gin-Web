package authcCenter

import (
	"fmt"
	mysqlDB "gin-web/initialize/mysql"
	"gin-web/utils"
	"strconv"
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
	UpdateTime time.Time `json:"updateTime" gorm:"column:updateTime;autoCreateTime"`        //修改time
	RoleID     int64     `json:"role_id" gorm:"column:role_id;type:bigint;not null"`        // 属于那个角色
	Role       Role      `gorm:"many2many:user_roles"`                                      //外键role
}

func (u *User) Add(user User) error {
	user.CreateTime = time.Now()
	if err := mysqlDB.DB.Model(&Role{}).Create(user).Error; err != nil {
		return err
	}
	return nil

}

func (u *User) Deleted(idString string) error {
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

func (u *User) Update(user User) error {
	//修改参数
	user.UpdateTime = time.Now()
	if err := mysqlDB.DB.Model(&Role{}).Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) GetAll(currPage, PageSize, startTime, endTime string) ([]User, error) {
	skip, limit, err := utils.GetPage(currPage, PageSize)
	if err != nil {
		return nil, err
	}
	tx := mysqlDB.DB
	if startTime != "" && endTime != "" {
		tx = tx.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var count int64
	var resDB []User
	res := tx.Model(&User{}).Limit(limit).Offset(skip).Find(&resDB).Count(&count)
	if res.Error != nil {
		return nil, err
	}
	return resDB, nil
}

func (u *User) IsExist(account string) (bool, error) {
	//查重
	var count int64
	err := mysqlDB.DB.Model(&User{}).Where("account = ?", account).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return true, nil
	}
	return false, fmt.Errorf("存在记录%d", count)
}

func (u *User) ChangePassword() {

}
