package authcCenter

import "time"

type Role struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Users      []User    `json:"users"`
	Apis       []Api     `json:"apis"`
	Desc       string    `json:"desc"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}
