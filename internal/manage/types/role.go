package types

import (
	"gin-web/internal/manage/models"
)

// 查询角色列表
type RoleGetListReq struct {
	Name      string `json:"name"`
	CurrPage  string `json:"currPage"`
	PageSize  string `json:"pageSize"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type RoleGetListResp struct {
	Roles []models.Role `json:"roles"`
	Total int64         `json:"total"`
}

// 新增角色
type RoleInsertReq struct {
	Name string `json:"name" binding:"required,min=2,max=20" `
	Desc string `json:"desc"  `
}

// 删除角色
type RoleDeleteReq struct {
	Id int64 `json:"id" binding:"required"`
}

// 修改角色
type RoleUpdateReq struct {
	Id   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required,min=2,max=20" `
	Desc string `json:"desc"  `
}
