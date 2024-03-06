package config

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"loopy-manager/initialize/global"
)

func MysqlInit(config MysqlConfig) {
	var err error
	global.MysqlClient, err = gorm.Open(mysql.Open(config.Address), &gorm.Config{})
	if err != nil {
		logrus.Fatalln("Mysql数据库连接失败:", err)
	}
	sqlDB, err := global.MysqlClient.DB()
	if err != nil {
		logrus.Fatalln("Mysql连接池创建失败")
	}
	sqlDB.SetMaxIdleConns(config.SetMaxIdleConns)       //最大空闲连接数
	sqlDB.SetMaxOpenConns(config.SetMaxOpenConns)       //最大连接数
	sqlDB.SetConnMaxLifetime(config.SetConnMaxLifetime) //设置连接空闲超时
	{
		global.UserTable = global.MysqlClient.Table("users")
		global.RoleTable = global.MysqlClient.Table("roles")
		global.ApiTable = global.MysqlClient.Table("apis")
		global.UserRoleTable = global.MysqlClient.Table("user_roles")
		global.RoleApiTable = global.MysqlClient.Table("role_apis")
		global.LogTable = global.MysqlClient.Table("operation_logs")
	}
	logrus.Println("mysql连接成功")
}
