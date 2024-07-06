package service

import (
	"fmt"
	"loopy-manager/app/global"
	"loopy-manager/app/model"
	"time"
)

//func GetRequestLog() {
//	var log []model.OperationLog
//	err := global.LogTable.Select("*,COUNT(*)").Group("method").Find(&log)
//	if err != nil {
//		fmt.Println("ERR", err)
//	}
//	fmt.Println("解码数据:", log)
//}

func GetRequestLog() {
	//status
	var logs []model.OperationLog
	err := global.LogTableSlave0.Debug().Select("method").Group("method").Find(&logs).Error
	if err != nil {
		fmt.Println("ERR", err)

	}
	fmt.Println("解码数据:", logs)
}

func GetRequestLog0() {
	var logs []model.OperationLog
	go func(mike []model.OperationLog) {
		err := global.MysqlClientSlave0.Debug().Where("status >(?)", global.LogTableSlave0.Select("AVG(status)")).Find(&logs).Error
		if err != nil {
			fmt.Println("ERR", err)
			return
		}
		for _, v := range mike {
			fmt.Println("当前status:", v.Status)
		}
	}(logs)
	time.Sleep(5 * time.Second)
	//err := global.MysqlClient.Debug().Where("status >(?)", global.LogTable.Select("AVG(status)")).Find(&logs).Error
	//if err != nil {
	//	fmt.Println("ERR", err)
	//	return
	//}
	//for _, v := range logs {
	//	fmt.Println("当前status:", v.Status)
	//}
}
