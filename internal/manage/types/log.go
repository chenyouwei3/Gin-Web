package types

import "gin-web/internal/manage/models"

// 查询操作日志列表
type LogByOperationGetListReq struct {
	Account   string `json:"account"`
	CurrPage  string `json:"currPage"`
	PageSize  string `json:"pageSize"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type LogByOperationGetListResp struct {
	Logs  []models.OperationLog `json:"logs"`
	Total int64                 `json:"total"`
}

// 查询运行日志
type LogByRunGetListReq struct {
	CurrPage  string `json:"currPage"`
	PageSize  string `json:"pageSize"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
