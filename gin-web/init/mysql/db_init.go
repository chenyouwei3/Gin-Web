package mysql

import (
	"fmt"
	conf "gin-web/init/config"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB(conf conf.MysqlConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
		conf.Charset,
	)
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
	sqlDB.SetMaxIdleConns(conf.SetMaxIdleConns)
	// 设置最大打开连接数为 100
	sqlDB.SetMaxOpenConns(conf.SetMaxOpenConns)
	// 设置连接最大生命周期为 30 秒
	sqlDB.SetConnMaxLifetime(conf.SetConnMaxLifetime)
	DB = db
	return nil
}
