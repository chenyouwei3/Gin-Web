package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"loopy-manager/initialize/config/system"
	"loopy-manager/initialize/global"
)

func MysqlInit(config system.MysqlConfig) {
	//global.MysqlClientMaster.AutoMigrate(model.OperationLog{}, model.User{}, model.Role{}, model.Api{})//初始化数据库
	mysqlInitMaster(config)
	mysqlInitSlave(config)
}

// mysql主从一致
func mysqlInitMaster(config system.MysqlConfig) {
	var err error
	global.MysqlClientMaster, err = gorm.Open(mysql.Open(config.Address0), &gorm.Config{})
	if err != nil {
		logrus.Fatalln("Mysql数据库连接失败:", err)
	}
	sqlDB, err := global.MysqlClientMaster.DB()
	if err != nil {
		logrus.Fatalln("Mysql连接池创建失败")
	}
	sqlDB.SetMaxIdleConns(config.SetMaxIdleConns)       //最大空闲连接数
	sqlDB.SetMaxOpenConns(config.SetMaxOpenConns)       //最大连接数
	sqlDB.SetConnMaxLifetime(config.SetConnMaxLifetime) //设置连接空闲超时
	{
		global.UserTableMaster = global.MysqlClientMaster.Table("users")
		global.RoleTableMaster = global.MysqlClientMaster.Table("roles")
		global.ApiTableMaster = global.MysqlClientMaster.Table("apis")
		global.UserRoleTableMaster = global.MysqlClientMaster.Table("user_roles")
		global.RoleApiTableMaster = global.MysqlClientMaster.Table("role_apis")
		global.LogTableMaster = global.MysqlClientMaster.Table("operation_logs")
	}
	logrus.Println("mysql-Master连接成功")
}

func mysqlInitSlave(config system.MysqlConfig) {
	var err error
	global.MysqlClientSlave0, err = gorm.Open(mysql.Open(config.Address0), &gorm.Config{})
	if err != nil {
		logrus.Fatalln("Mysql数据库连接失败:", err)
	}
	sqlDB, err := global.MysqlClientSlave0.DB()
	if err != nil {
		logrus.Fatalln("Mysql连接池创建失败")
	}
	sqlDB.SetMaxIdleConns(config.SetMaxIdleConns)       //最大空闲连接数
	sqlDB.SetMaxOpenConns(config.SetMaxOpenConns)       //最大连接数
	sqlDB.SetConnMaxLifetime(config.SetConnMaxLifetime) //设置连接空闲超时
	{
		global.UserTableSlave0 = global.MysqlClientSlave0.Table("users")
		global.RoleTableSlave0 = global.MysqlClientSlave0.Table("roles")
		global.ApiTableSlave0 = global.MysqlClientSlave0.Table("apis")
		global.UserRoleTableSlave0 = global.MysqlClientSlave0.Table("user_roles")
		global.RoleApiTableSlave0 = global.MysqlClientSlave0.Table("role_apis")
		global.LogTableSlave0 = global.MysqlClientSlave0.Table("operation_logs")
	}
	logrus.Println("mysql-Slave连接成功")
}
