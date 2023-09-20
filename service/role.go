package service

import (
	"LoopyTicker/global"
	"LoopyTicker/model"
	"LoopyTicker/utils"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func CreateRole(role model.Role) utils.Response {
	if role.Name == " " || role.Code == " " || len(role.Name) >= 20 || len(role.Code) >= 20 {
		return utils.ErrorMess("参数错误", role)
	}
	var roleDB []model.Role
	res := global.RoleTable.Where("name= ?", role.Name).Find(&roleDB)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return utils.ErrorMess("查重错误", res.Error.Error())
	}
	for _, v := range roleDB {
		if v.Name == role.Name {
			return utils.ErrorMess("role已存在", nil)
		}
	}
	role.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	role.Id = global.RoleSnowFlake.Generate().Int64()
	if len(role.Apis) != 0 {
		//json.Marshal()：将数据结构体struct转换为json字符串
		temp, err := json.Marshal(role.Apis)
		if err != nil {
			return utils.ErrorMess("失败", nil)
		}
		role.Api = string(temp)
	}
	res = global.RoleTable.Create(&role)
	if res.Error != nil {
		return utils.ErrorMess("失败1", res.Error.Error())
	}
	return utils.SuccessMess("成功", res.RowsAffected)
}

func DeletedRole(idString string) utils.Response {
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return utils.ErrorMess("参数错误", err.Error())
	}
	// 创建新的 *gorm.DB 对象
	//db := global.RoleTable.Session(&gorm.Session{})
	// 执行删除操作
	res := global.RoleTable.Session(&gorm.Session{}).Delete(&model.Role{Id: id})
	fmt.Println(res)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}
	return utils.SuccessMess("成功", res.RowsAffected)
}

func UpdateRole(role model.Role) utils.Response {
	if role.Id == 0 || role.Code == " " {
		return utils.ErrorMess("参数错误", nil)
	}
	var roleDB model.Role
	res := global.RoleTable.Session(&gorm.Session{}).Where("id = ?", role.Id).First(&roleDB)
	if res.Error != nil {
		return utils.ErrorMess("失败,该角色不存在", res.Error.Error())
	}
	role.CreateTime = roleDB.CreateTime
	role.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	temp, err := json.Marshal(role.Apis)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}
	role.Api = string(temp)
	res = global.RoleTable.Session(&gorm.Session{}).Where("id=?", role.Id).Save(&role)
	//res = global.RoleTable.Update(&role)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}
	return utils.SuccessMess("成功", res.RowsAffected)
}

func GetRole(name, currPage, pageSize, startTime, endTime string) utils.Response {
	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}

	if startTime != "" && endTime != "" {
		global.RoleTable = global.RoleTable.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var count int64
	var roleDB []model.Role
	res := global.RoleTable.Order("id desc").Where("name like ?", "%"+name+"%").Limit(limit).Offset(skip).Find(&roleDB).Count(&count)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}
	return utils.SuccessMess("成功", struct {
		Count int64        `json:"count" bson:"count"`
		Data  []model.Role `json:"data" bson:"data"`
	}{
		Count: count,
		Data:  roleDB,
	})
}
