package main

import (
	"loopy-manager/initialize"
	"loopy-manager/internal/router"
)

func init() {
	initialize.InitConfig()
}

func main() {
	//global.MysqlClientMaster.AutoMigrate(model.Moment{}, model.Comment{})
	engine := router.GetEngine()
	if err := engine.Run(":8098"); err != nil {
		panic(err)
	}
}
