package model

type Log struct {
	ID         int64       `db:"id" json:"id" gorm:"column:id;type:bigint;primary_key;not null"`
	User       interface{} `db:"user" json:"user" gorm:"column:user;type:varchar(225);"`
	Path       string      `db:"path" json:"path" gorm:"column:path;type:varchar(30);not null"`
	Method     string      `db:"method" json:"method" gorm:"column:method;type:varchar(10);not null"`
	Status     int         `db:"status" json:"status" gorm:"column:status;type:varchar(10);not null"`
	Query      string      `db:"query" json:"query" gorm:"column:query;type:varchar(30);not null"`
	Body       interface{} `db:"body" json:"body" gorm:"column:body;type:varchar(225);"`
	IP         string      `db:"ip" json:"ip" gorm:"column:ip;type:varchar;primary_key(30);not null"`
	UserAgent  string      `db:"userAgent" json:"userAgent" gorm:"column:userAgent;type:varchar(225);not null"`
	Errors     string      `db:"errors" json:"errors" gorm:"column:errors;type:varchar(225);"`
	Cost       string      `db:"cost" json:"cost" gorm:"column:cost;type:varchar(10);not null"`
	CreateTime string      `db:"createTime" json:"createTime" gorm:"column:createTime;type:varchar(20);not null"`
}
