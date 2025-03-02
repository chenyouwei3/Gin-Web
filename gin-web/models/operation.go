package models

import "time"

type OperationLog struct {
	Id        int64     `json:"id" gorm:"column:id;type:bigint;primaryKey;not null"`                       // 主键
	Username  string    `json:"username" gorm:"column:username;type:varchar(20);not null;comment:'用户登录名'"` // 用户名
	Ip        string    `json:"ip" gorm:"column:ip;type:varchar(20);comment:'IP地址'"`                       // IP 地址
	Method    string    `json:"method" gorm:"column:method;type:varchar(20);comment:'请求方式'"`               // 请求方式
	Status    int       `json:"status" gorm:"column:status;type:int(4);comment:'响应状态码'"`                   // 响应状态码
	Query     string    `json:"query" gorm:"column:query;type:varchar(50);comment:'查询内容'"`                 // 查询内容
	Body      string    `json:"body" gorm:"column:body;type:varchar(50);comment:'请求体'"`                    // 请求体
	Path      string    `json:"path" gorm:"column:path;type:varchar(100);not null;comment:'访问路径'"`         // 访问路径
	StartTime time.Time `json:"startTime" gorm:"column:start_time;type:datetime(3);comment:'发起时间'"`        // 发起时间
	UserAgent string    `json:"userAgent" gorm:"column:user_agent;type:varchar(50);comment:'浏览器标识'"`       // 浏览器标识
	CostTime  string    `json:"costTime" gorm:"column:cost_time;type:varchar(20);comment:'花费时间'"`          // 花费时间
}
