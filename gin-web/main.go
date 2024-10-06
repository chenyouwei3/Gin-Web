package main

import (
	"gin-web/initialize/config"
	mysqlDB "gin-web/initialize/mysql"
	"gin-web/models/authcCenter"
	"time"
)

func init() {
	config.InitConfig()
	mysqlDB.InitDB()
}

func main() {
	migration()
	time.Sleep(time.Hour)
	//gin.SetMode(gin.ReleaseMode)
}

func migration() {
	// 定义要迁移的表
	tables := []interface{}{
		&authcCenter.User{},
		&authcCenter.Role{},
		&authcCenter.Api{},
	}
	for _, table := range tables {
		if mysqlDB.DB.Migrator().HasTable(table) {
			return // 如果表存在，直接返回
		}
	}
	err := mysqlDB.DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&authcCenter.User{},
		&authcCenter.Role{},
		&authcCenter.Api{})
	if err != nil {
		panic(err)
	}
}
