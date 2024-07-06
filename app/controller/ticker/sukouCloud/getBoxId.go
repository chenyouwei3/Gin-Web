package sukonCloud

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"loopy-manager/app/global"
	"loopy-manager/app/model"
	"loopy-manager/pkg/redisUtils"
	"loopy-manager/pkg/utils"
	"net/url"
	"sync"
	"time"
)

func SuKonCloudProjects() { //获取项目
	URL := "http://sukon-cloud.com/api/v1/base/projects"
	urlValues := url.Values{}
	urlValues.Add("token", global.SukonCloudToken)
	var data model.SuKonProject
	data, err := utils.SukouCloudHTTP(data, URL, urlValues)
	if err != nil {
		logrus.Error("http请求失败:", err)
	}
	if data.Data == nil && data.Msg == "token已过期" {
		logrus.Error("project数据为空,获取失败---", data)
		utils.SukonToken()
		return
	}
	for _, project := range data.Data {
		if project.Id == "rKWw9LNBQYH" { //瑞通碳素项目id
			continue
		}
		suKonCloudBox(project.Id)
	}
}

func suKonCloudBox(projectId string) { //获取box并且更新box状态
	URL := "http://sukon-cloud.com/api/v1/base/projectBoxes"
	urlValues := url.Values{}
	urlValues.Add("token", global.SukonCloudToken)
	urlValues.Add("projectId", projectId)
	var data model.ProjectBox
	data, err := utils.SukouCloudHTTP(data, URL, urlValues)
	if err != nil {
		logrus.Error("http请求失败:", err)
	}
	if data.Success == false {
		logrus.Error("获取box异常", data)
	}
	var wg sync.WaitGroup
	for i, box := range data.Data {
		switch box.Status {
		case "0":
			var device model.Device
			strRedis, err := redisUtils.Redis{}.GetValueHash("rtts", box.BoxId)
			if err != nil {
				logrus.Error("解码失败:", err)
			}
			err = json.Unmarshal([]byte(strRedis), &device) //解码
			if err != nil {
				logrus.Error("解码失败:", err)
			}
			device.Status = "离线"
			device.UpdateTime = utils.TimeFormat(time.Now())
			lastDB, err := json.Marshal(device)
			if err != nil {
				logrus.Error("解码失败:", err)
			}
			err = redisUtils.Redis{}.SetValueHash("rtts", box.BoxId, string(lastDB))
			if err != nil {
				logrus.Error("解码失败:", err)
			}
			continue
		case "1":
			var device model.Device
			strRedis, err := redisUtils.Redis{}.GetValueHash("rtts", box.BoxId)
			if err != nil {
				logrus.Error("解码失败:", err)
			}
			err = json.Unmarshal([]byte(strRedis), &device) //解码
			if err != nil {
				logrus.Error("解码失败:", err)
			}
			device.Status = "在线"
			device.UpdateTime = utils.TimeFormat(time.Now())
			lastDB, err := json.Marshal(device)
			if err != nil {
				logrus.Error("解码失败:", err)
			}
			err = redisUtils.Redis{}.SetValueHash("rtts", box.BoxId, string(lastDB))
			if err != nil {
				logrus.Error("解码失败:", err)
			}
			wg.Add(1)
			go func(boxId string, i int) {
				defer wg.Done()
				BoxPlc(boxId)
			}(box.BoxId, i)
		default:
			logrus.Println("没有设备")
		}
	}
	wg.Wait()
}
