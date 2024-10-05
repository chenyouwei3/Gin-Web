package authcCenter

import "time"

type User struct {
	Id         int64     `json:"id" gorm:"column:id;type:bigint;primaryKey;not null"`
	Name       string    `json:"name" gorm:"column:name;type:varchar(20);not null"`         //用户名
	Account    string    `json:"account" gorm:"column:account;type:varchar(20);not null"`   //账号
	Password   string    `json:"password" gorm:"column:password;type:varchar(20);not null"` //密码
	AvatarUrl  string    `json:"avatarUrl" gorm:"column:avatarUrl;type:varchar(30)"`        //头像Url
	Sex        string    `json:"sex" gorm:"column:sex;type:varchar(3);not null"`            //性别
	Email      string    `json:"email" gorm:"column:email;type:varchar(20);not null"`       //邮箱
	Salt       string    `json:"salt" gorm:"column:salt;type:varchar(20);not null"`         //盐加密
	CreateTime time.Time `json:"createTime" gorm:"column:createTime;autoCreateTime"`        //创建time
	UpdateTime time.Time `json:"updateTime" gorm:"column:updateTime;autoCreateTime"`        //修改time
	RoleID     int64     `json:"role_id" gorm:"column:role_id;type:bigint;not null"`        // 属于那个角色
	Role       Role      `gorm:"foreignKey:RoleID"`                                         //外键role
}

func (u *User) Add() {

}

func (u *User) Deleted() {

}

func (u *User) Update() {

}

func (u *User) GetAll() {

}

func (u *User) GetOne() {

}
