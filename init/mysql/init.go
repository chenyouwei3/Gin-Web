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
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
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
		DSN:                       dns,   // 数据源名称：格式为 "user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		DefaultStringSize:         256,   // string 类型字段默认长度为 256（用于创建表结构时自动生成 VARCHAR(256)）
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，兼容 MySQL 5.6（它不支持 datetime(3) 这种精度写法）
		DontSupportRenameIndex:    true,  // 禁止重命名索引（对不支持 RENAME INDEX 的 MySQL 版本，GORM 会转为 Drop + Create 实现）
		DontSupportRenameColumn:   true,  // 禁止重命名字段（对不支持 RENAME COLUMN 的版本，使用删除 + 添加字段）
		SkipInitializeWithVersion: false, // 是否跳过获取 MySQL 版本的自动初始化，设置为 false 则默认去检测版本
	}), &gorm.Config{
		Logger: ormLogger, // 指定自定义的 GORM 日志器（可用于控制日志级别、格式等）
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，例如 `User` 表名对应 `user` 而不是默认的 `users`
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
