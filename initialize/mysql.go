package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"loopy-manager/global"
)

const dsn = "root:Cyw123456@tcp(43.138.32.203:3306)/loopyticker?charset=utf8mb4&parseTime=True&loc=Local"

// MysqlInit 初始化mysql
func MysqlInit() {
	var err error
	global.MysqlClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Mysql数据库连接失败%s", err)
	}
	fmt.Println("mysql连接成功")
	//设置连接池数量
	//sqlDB, err := global.MysqlClient.DB()
	//if err != nil {
	//	fmt.Println("连接池失败")
	//}
	//sqlDB.SetMaxIdleConns(10)                  //最大空闲连接数
	//sqlDB.SetMaxOpenConns(10)                  //最大连接数
	//sqlDB.SetConnMaxLifetime(time.Minute * 15) //设置连接空闲超时
	{
		global.UserTable = global.MysqlClient.Table("user")
		global.RoleTable = global.MysqlClient.Table("role")
		global.ApiTable = global.MysqlClient.Table("api")
		global.LogTable = global.MysqlClient.Table("log")
	}
}
