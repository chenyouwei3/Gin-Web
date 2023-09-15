package initialize

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"loopy-manager/global"
)

const MysqlUrl = "admin:admin@tcp(127.0.0.1:3306)/sdlManager"

func MysqlInit() {
	var err error
	if global.MysqlClient == nil {
		global.MysqlClient, err = gorm.Open("mysql", MysqlUrl)
		if err != nil {
			fmt.Println("mysql连接失败", err)
		}
	}
}
