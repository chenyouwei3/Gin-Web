package main

import (
	conf "gin-web/init/config"
	mysqlDB "gin-web/init/mysql"
	"gin-web/internal/manage/models"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	t  time.Time
)

func init() {
	//配置文件加载
	err := conf.InitConfig("../configs", "config-app1", "yml")
	if err != nil {
		log.Println("配置文件加载失败", err)
		panic(err)
	}
	//设置gin框架运行模式
	if conf.SystemConfig.APP.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	}
	//设置mysql数据库
	if err = mysqlDB.InitDB(conf.SystemConfig.Mysql); err != nil {
		log.Println("MySql初始化失败", err)
		panic(err)
	}

	log.Println("系统初始化加载完毕")
	db = mysqlDB.DB
	if err = mysqlDB.DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.OperationLog{},
	); err != nil {
		panic(err)
	}
}

func main() {
	var err error
	role_admin := &models.Role{
		1,
		time.Now(),
		t,
		"super admin",
		""}
	//role
	role_user := models.Role{
		2,
		time.Now(),
		t,
		"normal user",
		""}
	if err = role_user.Insert(); err != nil {
		panic(err)
	}
	if err = role_admin.Insert(); err != nil {
		panic(err)
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte("123456"), 12)
	if err != nil {
		panic(err)
	}
	user := &models.User{
		1,
		time.Now(),
		t,
		"chenyouwei",
		"chenyouwei3@outlook.com",
		"21480",
		string(bytes),
		"",
		nil,
	}
	if err = user.Insert([]int{1}); err != nil {
		panic(err)
	}
}
