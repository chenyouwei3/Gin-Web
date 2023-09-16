package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Account   string             `bson:"account" json:"account"` //账号
	Password  string             `bson:"password,omitempty" json:"password,omitempty"`
	Name      string             `bson:"name" json:"name"`           //用户名
	AvatarUrl string             `bson:"avatarUrl" json:"avatarUrl"` //头像地址
	Sex       string             `bson:"sex" json:"sex"`
	Phone     string             `bson:"phone" json:"phone"`
	Salt      string             `bson:"salt,omitempty" json:"salt,omitempty"`
	RoleId    primitive.ObjectID `bson:"roleId" json:"roleId"`
	OpenId    string             `bson:"openId" json:"openId"`
}
