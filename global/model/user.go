package model

type User struct {
	Id         int64   `db:"id" json:"id" gorm:"column:id;type:bigint;primary_key;not null"`
	Name       string  `db:"name" json:"name" gorm:"column:name;type:varchar(20);not null"`
	Account    string  `db:"account" json:"account" gorm:"column:account;type:varchar(20);not null"`
	Password   string  `db:"password" json:"password" gorm:"column:password;type:varchar(60);not null"`
	AvatarUrl  string  `db:"avatarUrl" json:"avatarUrl" gorm:"column:avatarUrl;type:varchar(144);not null"` //头像地址
	IsValid    bool    `db:"isValid" json:"isValid" gorm:"column:isValid;type:tinyint(1);not null;default:1"`
	Sex        string  `db:"sex" json:"sex" gorm:"column:sex;type:varchar(2);not null"`
	Phone      string  `db:"phone" json:"phone" gorm:"column:phone;type:varchar(11);not null"`
	Mail       string  `db:"mail" json:"mail" gorm:"column:mail;type:varchar(11);not null"`
	Salt       string  `db:"salt" json:"salt" gorm:"column:salt;type:varchar(60);not null"`
	RoleId     string  `db:"roleId" json:"roleId" gorm:"column:roleId;type:text"`
	RoleIds    []int64 `db:"roleIds" json:"roleIds" gorm:"-"`
	CreateTime string  `db:"createTime" json:"createTime" gorm:"column:createTime;type:varchar(20);not null"`
	UpdateTime string  `db:"updateTime" json:"updateTime" gorm:"column:updateTime;type:varchar(20)"`
}
