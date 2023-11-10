package initialize

import "loopy-manager/utils"

func RabbitMqSendInit() {
	//rabbitSimple := utils.NewRabbitConn("simple", "", "")
	//rabbitSimple.PublishSimple("simple")
	//rabbitWork := utils.NewRabbitConn("work", "", "")
	//for {
	//	rabbitWork.PublishWork("work")
	//	fmt.Println("发送成功")
	//	time.Sleep(1 * time.Second)
	//}
	rabbitPublish := utils.NewRabbitConn("publish", "logs", "")
	rabbitPublish.PublishPublish("666")
}
