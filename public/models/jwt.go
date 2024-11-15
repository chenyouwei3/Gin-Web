package models

import "time"

type User struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`        //用户名
	Account    string    `json:"account"`     //账号
	Password   string    `json:"password"`    //密码
	AvatarUrl  string    `json:"avatarUrl"`   //头像Url
	Sex        string    `json:"sex"`         //性别
	Email      string    `json:"email"`       //邮箱
	Salt       string    `json:"salt"`        //盐加密
	CreateTime time.Time `json:"createTime"`  //创建time
	UpdateTime time.Time `json:"updateTime" ` //修改time
}
