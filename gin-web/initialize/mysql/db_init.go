package mysql

import (
	"fmt"
	"gin-web/initialize/config"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB

func InitDB() error {
	confDB := config.Conf.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
		confDB.UserName, confDB.Password, confDB.Host, confDB.Port, confDB.Database, confDB.Charset)
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	sqlDB, _ := db.DB()
	// 设置最大空闲连接数为 20
	sqlDB.SetMaxIdleConns(20)
	// 设置最大打开连接数为 100
	sqlDB.SetMaxOpenConns(100)
	// 设置连接最大生命周期为 30 秒
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db
	return nil
}
