package service

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"loopy-manager/global"
	"loopy-manager/middleware"
	"loopy-manager/model"
	"loopy-manager/utils"
	"math/rand"
	"strconv"
	"time"
)

func Login(user model.User) utils.Response {
	if user.Account == "" || user.Password == "" {
		return utils.ErrorMess("登录失败", nil)
	}
	var userDB model.User
	res := global.UserTable.Where("account=?", user.Account).Find(&userDB)
	if res.Error != nil {
		return utils.ErrorMess("该用户不存在", nil)
	}
	//密码验证
	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password+userDB.Salt)); err != nil {
		return utils.ErrorMess("密码错误", err)
	}
	//json反序列化
	err := json.Unmarshal([]byte(userDB.RoleId), &userDB.RoleIds)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}
	token, err := middleware.CreateToken(userDB)
	if err != nil {
		return utils.ErrorMess("生成token失败", err.Error())
	}
	data := map[string]interface{}{
		"id":      userDB.Id,
		"name":    userDB.Name,
		"account": userDB.Account,
		"roleId":  userDB.RoleId,
		"roleIds": userDB.RoleIds,
		"sex":     userDB.Sex,
		"token":   token,
	}
	return utils.SuccessMess("成功", data)
}

func Register(user model.User) utils.Response {
	//查询
	if user.Account == "" || user.Password == "" {
		return utils.ErrorMess("账号密码不能为空", nil)
	}
	var userDB []model.User
	res := global.UserTable.Where("account=?", user.Account).Find(&userDB)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error)
	}
	//查看是否查出来有数据
	if res.RowsAffected != 0 {
		return utils.ErrorMess("该用户已存在", nil)
	}
	//加密
	//根据时间戳生成种子，防止恶意伪造
	rand.New(rand.NewSource(time.Now().Unix()))
	//生成盐
	user.Salt = strconv.FormatInt(time.Now().Unix(), 10)
	//密码加密加盐
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password+user.Salt), bcrypt.DefaultCost)
	if err != nil {
		return utils.ErrorMess("密码加密失败", err.Error())
	}
	user.Password = string(encryptedPassword)
	user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	//x雪花算法生成分布式id
	user.Id = global.UserSnowFlake.Generate().Int64()
	if len(user.RoleIds) != 0 {
		//json.Marshal()：将数据结构体struct转换为json字符串
		fmt.Println(user.RoleId)
		temp, err := json.Marshal(user.RoleIds)
		if err != nil {
			return utils.ErrorMess("失败", nil)
		}
		user.RoleId = string(temp)
	}
	res = global.UserTable.Create(&user)
	if res.Error != nil {
		return utils.ErrorMess("创建失败", res.Error.Error())
	}
	return utils.SuccessMess("成功", res.RowsAffected)
}
