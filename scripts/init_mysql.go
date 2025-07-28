package main

import (
	conf "gin-web/init/config"
	mysqlDB "gin-web/init/mysql"
	"gin-web/internal/manage/models"
	"github.com/gin-gonic/gin"
	"log"
	"time"

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
	////api
	//apis := []authCenter.Api{
	//	{1, "增加api", "/api/insert", "POST", "系统初始化", time.Now(), t},
	//	{2, "删除api", "/api/remove", "POST", "系统初始化", time.Now(), t},
	//	{3, "更新api", "/api/edit", "POST", "系统初始化", time.Now(), t},
	//	{4, "查询api", "/api/getList", "GET", "系统初始化", time.Now(), t},
	//
	//	{5, "增加role", "/role/insert", "POST", "系统初始化", time.Now(), t},
	//	{6, "删除role", "/role/remove", "POST", "系统初始化", time.Now(), t},
	//	{7, "更新role", "/role/edit", "POST", "系统初始化", time.Now(), t},
	//	{8, "查询role", "/role/getList", "GET", "系统初始化", time.Now(), t},
	//	{9, "根据role_id查api", "/role/getRoleByApis", "GET", "系统初始化", time.Now(), t},
	//
	//	{10, "增加user", "/user/insert", "POST", "系统初始化", time.Now(), t},
	//	{11, "删除user", "/user/remove", "POST", "系统初始化", time.Now(), t},
	//	{12, "更新user", "/user/edit", "POST", "系统初始化", time.Now(), t},
	//	{13, "查询user", "/user/getList", "GET", "系统初始化", time.Now(), t},
	//	{14, "根据user_id查role", "/user/getUserByRoles", "GET", "系统初始化", time.Now(), t},
	//
	//	{15, "查询操作日志", "/log/operation/getList", "GET", "系统初始化", time.Now(), t},
	//
	//	{16, "文件列表", "/dist/list", "GET", "系统初始化", time.Now(), t},
	//	{17, "创建文件夹", "/dist/mkdir", "POST", "系统初始化", time.Now(), t},
	//	{18, "重命名(文件/文件夹)", "/dist/rename", "POST", "系统初始化", time.Now(), t},
	//	{19, "删除(文件/文件夹)", "/dist/remove", "POST", "系统初始化", time.Now(), t},
	//	{20, "复制(文件/文件夹)", "/dist/copy", "POST", "系统初始化", time.Now(), t},
	//	{21, "移动(文件/文件夹)", "/dist/move", "POST", "系统初始化", time.Now(), t},
	//	{22, "移动(文件/文件夹)下拉框", "/dist/dropdownMenu", "GET", "系统初始化", time.Now(), t},
	//	{23, "下载(文件/文件夹)*", "/dist/download", "POST", "系统初始化", time.Now(), t},
	//	{24, "上传文件*", "/dist/upload", "POST", "系统初始化", time.Now(), t},
	//	{25, "文件可视化数据统计", "/dist/category", "GET", "系统初始化", time.Now(), t},
	//	{26, "文件在线预览文件", "/dist/onlicePreview", "GET", "系统初始化", time.Now(), t},
	//}
	//if err := db.Create(&apis).Error; err != nil {
	//	panic(err)
	//}
	////role
	//role_superAdmin := authCenter.Role{1, "超级管理员", "系统初始化", time.Now(), t, nil}
	//role_superAdmin.Insert([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26})
	//
	//role_normalAdmin := authCenter.Role{2, "管理员", "系统初始化", time.Now(), t, nil}
	//role_normalAdmin.Insert([]int{13, 15})
	//
	//role_user := authCenter.Role{3, "普通用户", "系统初始化", time.Now(), t, nil}
	//if err := role_user.Insert([]int{16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}); err != nil {
	//	panic(err)
	//}
	//
	////user
	//user_init := authCenter.User{1, "超级管理员", "super_admin", "$2a$10$IUcsDNBFQaF3vOm7hAfADOFJQRoZE02k5lCsXXuse2vA1vbLG/zWS", "", "女", "test@test.com", "9221363090780876532", time.Now(), t, nil}
	//if err := user_init.Insert([]int{1, 2, 3}); err != nil {
	//	panic(err)
	//}
}
