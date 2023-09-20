package service

import (
	"LoopyTicker/global"
	"LoopyTicker/model"
	"LoopyTicker/utils"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func CreateApi(api model.Api) utils.Response {
	if api.Name == " " ||
		api.Url == " " ||
		len(api.Url) >= 20 ||
		len(api.Name) >= 10 ||
		(api.Method != "GET" && api.Method != "POST" && api.Method != "PUT" && api.Method != "DELETE") {
		return utils.ErrorMess("参数错误", nil)
	}
	var apiDB []model.Api
	res := global.ApiTable.Where("name = ?", api.Name).Or("url = ? and method = ?", api.Url, api.Method).Find(&apiDB)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return utils.ErrorMess("查重错误", res.Error.Error())
	}
	for _, v := range apiDB {
		if v.Url == api.Url || v.Name == api.Name {
			return utils.ErrorMess("api已存在", nil)
		}
	}
	api.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	api.Id = global.ApiSnowFlake.Generate().Int64()
	res = global.ApiTable.Create(&api)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}
	return utils.SuccessMess("成功", res.RowsAffected)
}

func DeletedApi(idString string) utils.Response {
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}
	res := global.ApiTable.Session(&gorm.Session{}).Delete(&model.Api{Id: id})
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}
	return utils.SuccessMess("成功", res.RowsAffected)
}

func UpdateApi(api model.Api) utils.Response {
	if api.Id == 0 ||
		api.Name == "" ||
		api.Url == "" ||
		len(api.Url) >= 20 ||
		len(api.Name) >= 10 ||
		(api.Method != "GET" && api.Method != "POST" && api.Method != "PUT" && api.Method != "DELETE") {
		return utils.ErrorMess("失败,参数错误", nil)
	}
	var apiDB model.Api
	res := global.ApiTable.Session(&gorm.Session{}).Where("id=?", api.Id).First(&apiDB)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return utils.ErrorMess("失败,该API不存在", nil)
	}
	apiDB = api
	apiDB.CreateTime = api.CreateTime
	apiDB.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	res = global.ApiTable.Session(&gorm.Session{}).Where("id = ?", api.Id).Save(&apiDB)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}
	return utils.SuccessMess("成功", res.RowsAffected)
}

func GetApi(name, currPage, pageSize, startTime, endTime string) utils.Response {
	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}
	if startTime != "" && endTime != "" {
		global.ApiTable = global.ApiTable.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var count int64
	var apiDB []model.Api
	res := global.ApiTable.Order("id desc").Where("name like ?", "%"+name+"%").Limit(limit).Offset(skip).Find(&apiDB).Count(&count)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}
	return utils.SuccessMess("成功", struct {
		Count int64       `json:"count" bson:"count"`
		Data  []model.Api `json:"data" bson:"data"`
	}{
		Count: count,
		Data:  apiDB,
	})
}
