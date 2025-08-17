package types

import (
	"gin-web/internal/manage/models"
)

// 查询用户列表
type UserGetListReq struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CurrPage  string `json:"currPage"`
	PageSize  string `json:"pageSize"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type UserGetListResp struct {
	Users []models.User `json:"users"`
	Total int64         `json:"total"`
}

// 新增用户
type UserInsertReq struct {
	User struct {
		Name      string `json:"name" binding:"required,min=2,max=35"`       /// 名字必填，2~35长度
		Email     string `json:"email" binding:"required,email"`             // 必填且为邮箱格式
		Account   string `json:"account" binding:"required,min=4,max=35"`    // 账号必填
		Password  string `json:"password" binding:"required,min=6,max=100" ` // 密码必填且长度约束
		AvatarUrl string `json:"avatarUrl" binding:"omitempty,url"`          // 头像URL可选但格式必须是URL
	} `json:"user"`
	AddRoles     []int `json:"addRoles"`
	DeletedRoles []int `json:"deletedRoles"`
}

// 删除用户
type UserDeleteReq struct {
	Id int64 `json:"id" binding:"required"`
}

// 更新用户
type UserUpdateReq struct {
	User struct {
		Id        int64  `json:"id" binding:"required"`
		Name      string `json:"name" binding:"required,min=2,max=35"` /// 名字必填，2~35长度
		Email     string `json:"email" binding:"required,email"`       // 必填且为邮箱格式
		Account   string `json:"account"`                              // 账号必填
		Password  string `json:"password"`                             // 密码必填且长度约束
		AvatarUrl string `json:"avatarUrl"`                            // 头像URL可选但格式必须是URL
	} `json:"user"`
	AddRoles     []int `json:"addRoles"`
	DeletedRoles []int `json:"deletedRoles"`
}

// 用户登录
type UserLoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	User         *models.User `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
}
