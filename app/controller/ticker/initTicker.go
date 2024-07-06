package ticker

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"loopy-manager/app/controller/ticker/sukouCloud"
	"loopy-manager/app/global"
	"loopy-manager/pkg/utils"
)

func CornTicker() {
	if utils.IsProd() {
		return
	}
	//utils.SukonToken() //更新全局变量SuKon-Token
	//sukonCloud.SuKonCloudProjects()
	c := cron.New() //新建一个定时任务对象
	//定时获取token
	addCornFunc(c, global.Spec, utils.SukonToken, "获取SukouCloud-Token失败")
	addCornFunc(c, "0 */1 * * * *", sukonCloud.SuKonCloudProjects, "开始SukouCloud项目:") //每分钟存储生产工艺数据
	c.Start()
	select {}
}

func addCornFunc(c *cron.Cron, spec string, cmd func(), str string) {
	err := c.AddFunc(spec, cmd)
	if err != nil {
		logrus.Warnln(str+":", err)
	}
}
