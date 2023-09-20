package model

type Api struct {
	Id         int64  `db:"id" json:"id" gorm:"column:id;type:bigint(20);primaryKey;not null"`
	Name       string `db:"name" json:"name" gorm:"column:name;type:varchar(20);not null"`
	Url        string `db:"url" json:"url" gorm:"column:url;type:varchar(20);not null"`
	Method     string `db:"method" json:"method" gorm:"column:method;type:varchar(10);not null"`
	Desc       string `db:"desc" json:"desc" gorm:"column:desc;type:varchar(144)"`
	CreateTime string `db:"createTime" json:"createTime" gorm:"column:createTime;type:varchar(20);not null"`
	UpdateTime string `db:"updateTime" json:"updateTime" gorm:"column:updateTime;type:varchar(20)"`
}
