package global

import "gorm.io/gorm"

var (
	MysqlClient *gorm.DB
	UserTable   *gorm.DB
	ApiTable    *gorm.DB
	RoleTable   *gorm.DB
	LogTable    *gorm.DB
)
