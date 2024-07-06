package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	global2 "loopy-manager/app/global"
	"loopy-manager/app/model"
	"loopy-manager/pkg/utils"
	"math/rand"
	"strconv"
	"time"
)

type UserService struct{}

func (U UserService) CreateUser(user model.User) utils.Response {
	if err := global2.UserTableMaster.Transaction(func(tx *gorm.DB) error {
		//查询账号重复
		var userDB model.User
		if err := tx.Debug().Select("account").Where("account = ?", user.Account).Find(&userDB).Error; (err != nil && !errors.Is(err, gorm.ErrRecordNotFound)) || userDB.Account == user.Account {
			return fmt.Errorf("账号已存在:%w", err)
		}
		// 查询角色是否存在
		var roleDB model.Role
		if err := global2.RoleTableMaster.Debug().Select("id").Where("id = ?", user.RoleID).Find(&roleDB).Error; err != nil {
			return fmt.Errorf("查询角色错误:%w", err)
		}
		rand.New(rand.NewSource(time.Now().Unix())) //根据时间戳生成种子
		salt := strconv.FormatInt(rand.Int63(), 10) //生成盐
		encryptedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password+salt), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("密码加密失败:%w", err)
		}
		user.Password, user.Salt = string(encryptedPass), salt
		//插入事务
		user.Id = global2.UserSnowFlake.Generate().Int64()
		if err := tx.Debug().Create(&user).Error; err != nil {
			return fmt.Errorf("创建用户失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("插入成功", "1")
}

func (U UserService) DeletedUser(idString string) utils.Response {
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}
	if err := global2.UserTableMaster.Transaction(func(tx *gorm.DB) error {
		// 删除用户记录
		if err := tx.Debug().Delete(&model.User{}, id).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("删除事务失败", err.Error())
	}
	return utils.SuccessMess("删除成功", id)
}

func (U UserService) UpdatedUser(user model.User) utils.Response {
	if err := global2.UserTableMaster.Transaction(func(tx *gorm.DB) error {
		var userDB model.User
		if err := tx.Where("id = ?", user.Id).Take(&userDB).Error; err != nil {
			return fmt.Errorf("查询失败%w", err)
		}
		userDB.Name = user.Name
		userDB.RoleID = user.RoleID
		if err := tx.Debug().Save(&userDB).Error; err != nil {
			return fmt.Errorf("更新角色失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("修改用户成功", user.Id)
}

func (U UserService) GetUser(name, currPage, pageSize, startTime, endTime string) utils.Response {
	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		return utils.ErrorMess("数据转化失败", err.Error())
	}
	tx := global2.UserTableSlave.ChooseSlave()
	if startTime != "" && endTime != "" {
		tx = tx.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var count int64
	var userDB []model.User
	//加载role的信息.Preload("Role")
	res := tx.Where("name like ?", "%"+name+"%").Limit(limit).Offset(skip).Find(&userDB).Count(&count)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}
	return utils.SuccessMess("成功", struct {
		Count int64        `json:"count" bson:"count"`
		Data  []model.User `json:"data" bson:"data"`
	}{
		Count: count,
		Data:  userDB,
	})
}
