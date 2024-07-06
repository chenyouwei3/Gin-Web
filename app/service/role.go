package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	global2 "loopy-manager/app/global"
	"loopy-manager/app/model"
	"loopy-manager/pkg/utils"
	"strconv"
)

type RoleService struct{}

func (R RoleService) CreateRole(role model.Role) utils.Response {
	if err := global2.RoleTableMaster.Transaction(func(tx *gorm.DB) error {
		//查询角色重复
		var roleDB model.Role
		if err := tx.Debug().Select("name").Where("name = ?", roleDB.Name).Take(&roleDB).Error; (err != nil && !errors.Is(err, gorm.ErrRecordNotFound)) || role.Name == roleDB.Name {
			return fmt.Errorf("角色重复:%w", err)
		}
		// 查询api是否存在
		var apiDB []model.Api
		if err := global2.ApiTableMaster.Select("id").Where("id IN ?", extractRoleID(role.Api)).Find(&apiDB).Error; err != nil {
			return fmt.Errorf("查询api错误:%w", err)
		}
		if len(apiDB) != len(role.Api) { // 检查查询到的api数量是否和传入的api数量相等
			return fmt.Errorf("api数量不相等")
		}
		//插入事务
		role.Id = global2.RoleSnowFlake.Generate().Int64()
		if err := tx.Debug().Create(&role).Error; err != nil {
			return fmt.Errorf("创建角色失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("插入成功", "role")
}

func extractRoleID(apis []model.Api) []int64 { // 提取角色ID列表(辅助函数)
	ids := make([]int64, len(apis))
	for i, api := range apis {
		ids[i] = api.Id
	}
	return ids
}

func (R RoleService) DeletedRole(idString string) utils.Response {
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return utils.ErrorMess("字符串转化整数失败", err.Error())
	}
	if err := global2.RoleTableMaster.Transaction(func(tx *gorm.DB) error {
		tx0 := global2.RoleApiTableMaster.Begin()
		if err := tx0.Model(&model.Role{Id: id}).Association("Api").Clear(); err != nil {
			tx0.Rollback()
			return fmt.Errorf("清除关联失败:%w", err)
		}
		tx0.Commit()
		// 删除角色记录
		if err := tx.Delete(&model.Role{}, id).Error; err != nil {
			return fmt.Errorf("删除角色失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("删除事务失败", err.Error())
	}
	return utils.SuccessMess("删除成功", id)
}

func (R RoleService) UpdateRole(role model.Role) utils.Response {
	if err := global2.RoleTableMaster.Transaction(func(tx *gorm.DB) error {
		var roleDB model.Role
		if err := tx.Where("id = ?", role.Id).First(&roleDB).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("查询失败%w", err)
		}
		roleDB.Name = role.Name
		roleDB.Api = role.Api
		//tx0 := global.RoleApiTable.Begin()
		//if err := tx0.Save(&roleDB.Api).Error; err != nil {
		//	tx0.Rollback()
		//	return fmt.Errorf("修改关联表失败:%w", err)
		//}
		//tx0.Commit()
		if err := tx.Save(&roleDB).Error; err != nil {
			return fmt.Errorf("更新角色失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("修改角色成功", role.Id)
}

func (R RoleService) GetRole(name, currPage, pageSize, startTime, endTime string) utils.Response {
	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		return utils.ErrorMess("数据转化失败", err.Error())
	}
	tx := global2.RoleTableSlave.ChooseSlave()
	if startTime != "" && endTime != "" {
		tx = tx.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var count int64
	var roleDB []model.Role
	//Order("id desc")id降序排列
	res := tx.Debug().Where("name like ?", "%"+name+"%").Limit(limit).Offset(skip).Find(&roleDB).Count(&count)
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
