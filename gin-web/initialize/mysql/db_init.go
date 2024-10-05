package mysql

import (
	"gin-web/initialize/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

var DB *gorm.DB

func InitDB() {
	host := viper.GetString(config.Conf.MySQL.Host)
	port := viper.GetString(config.Conf.MySQL.Port)
	database := viper.GetString(config.Conf.MySQL.Database)
	username := viper.GetString(config.Conf.MySQL.UserName)
	password := viper.GetString(config.Conf.MySQL.Password)
	charset := viper.GetString(config.Conf.MySQL.Charset)
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=" + charset + "&parseTime=true"}, "")
	err := databaseInit(dsn)
	if err != nil {
		panic(err)
	}
}

func databaseInit(connString string) error {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connString,
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
		panic(err)
	}
	sqlDB, _ := db.DB()
	// 设置最大空闲连接数为 20
	sqlDB.SetMaxIdleConns(20)
	// 设置最大打开连接数为 100
	sqlDB.SetMaxOpenConns(100)
	// 设置连接最大生命周期为 30 秒
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
	return err
}
