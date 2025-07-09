package models

import (
	mysqlDB "gin-web/init/mysql"
	"time"
)

type OperationLog struct {
	Id        int64     `json:"id" gorm:"column:id;type:bigint;primaryKey;not null"`                           // 主键
	Account   string    `json:"account" gorm:"column:account;type:varchar(20);not null;comment:'用户登录名';index"` // 用户名
	Ip        string    `json:"ip" gorm:"column:ip;type:varchar(20);comment:'IP地址'"`                           // IP 地址
	Method    string    `json:"method" gorm:"column:method;type:varchar(20);comment:'请求方式'"`                   // 请求方式
	Status    int       `json:"status" gorm:"column:status;type:int(4);comment:'响应状态码'"`                       // 响应状态码
	Query     string    `json:"query" gorm:"column:query;type:varchar(50);comment:'查询内容'"`                     // 查询内容
	Body      string    `json:"body" gorm:"column:body;type:varchar(225);comment:'请求体'"`                       // 请求体
	Path      string    `json:"path" gorm:"column:path;type:varchar(50);not null;comment:'访问路径'"`              // 访问路径
	StartTime time.Time `json:"startTime" gorm:"column:start_time;type:datetime(3);comment:'发起时间';index"`      // 发起时间，添加索引
	UserAgent string    `json:"userAgent" gorm:"column:user_agent;type:varchar(220);comment:'浏览器标识'"`          // 浏览器标识
	CostTime  string    `json:"costTime" gorm:"column:cost_time;type:varchar(20);comment:'花费时间'"`              // 花费时间
}

func (OperationLog) tableName() string {
	return "log_operation"
}

// 查询
func (o *OperationLog) GetList(skip, limit int, startTime, endTime string) ([]OperationLog, int64, error) {
	// 1. 先统计总数（不做分页）
	var total int64
	countTx := mysqlDB.DB.Model(&OperationLog{})
	if startTime != "" && endTime != "" {
		countTx = countTx.Where("start_time >= ? AND start_time <= ?", startTime, endTime)
	}
	if o.Account != "" {
		countTx = countTx.Where("account = ?", o.Account)
	}
	if err := countTx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 2. 然后子查询，拿这一页要显示的 id 列表
	subQuery := mysqlDB.DB.Model(&OperationLog{}).Select("id").Order("start_time DESC")
	if startTime != "" && endTime != "" {
		subQuery = subQuery.Where("start_time >= ? AND start_time <= ?", startTime, endTime)
	}
	if o.Account != "" {
		subQuery = subQuery.Where("account = ?", o.Account)
	}
	subQuery = subQuery.Offset(skip).Limit(limit)
	// 3. 主查询：回表查完整记录
	var resDB []OperationLog
	if err := mysqlDB.DB.Model(&OperationLog{}).
		Select("account", "ip", "method", "status", "query", "body", "path", "start_time", "user_agent", "cost_time").
		Joins("JOIN (?) AS tmp ON tmp.id = operation_log.id", subQuery).
		Order("start_time DESC").
		Find(&resDB).Error; err != nil {
		return nil, 0, err
	}
	return resDB, total, nil
}

/* func (o *OperationLog) GetAll(skip, limit int, startTime, endTime string) ([]OperationLog, int64, error) {
	//tx := mysqlDB.DB.Model(&OperationLog{}).
	//	Select("username", "ip", "method", "status", "query", "body", "path", "start_time", "user_agent", "cost_time").
	//	Order("start_time DESC") //DESC/ASC
	//if startTime != "" && endTime != "" {
	//	tx = tx.Where("start_time >= ? and start_time <=?", startTime, endTime)
	//}
	//if o.Username != "" {
	//	tx = tx.Where("username = ?", o.Username)
	//}
	//var resDB []OperationLog
	//var counts int64
	//res := tx.Limit(limit).Offset(skip).Find(&resDB).Count(&counts)
	//fmt.Println("testing", counts)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	//return resDB, nil
	// 第一步：子查询，仅查出需要分页的 id 列表
	subQuery := mysqlDB.DB.Model(&OperationLog{}).Select("id").Order("start_time DESC")
	var total int64
	if startTime != "" && endTime != "" {
		subQuery = subQuery.Where("start_time >= ? AND start_time <= ?", startTime, endTime)
	}
	if o.Username != "" {
		subQuery = subQuery.Where("username = ?", o.Username)
	}
	subQuery = subQuery.Offset(skip).Limit(limit).Count(&total)
	// 主查询：回表查完整记录
	var resDB []OperationLog
	err := mysqlDB.DB.Model(&OperationLog{}).
		Select("username", "ip", "method", "status", "query", "body", "path", "start_time", "user_agent", "cost_time").
		Joins("JOIN (?) AS tmp ON tmp.id = operation_log.id", subQuery).
		Order("start_time DESC").
		Find(&resDB).Error

	if err != nil {
		return nil, 0, err
	}
	return resDB, total, nil
} */
