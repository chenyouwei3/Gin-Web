package config

import "time"

type Config struct {
	APP struct {
		Mode string //主程序运行模式
	}
	Mysql MysqlConfig
}

type MysqlConfig struct {
	Host               string
	Port               string
	Username           string
	Password           string
	Database           string
	Charset            string
	SetMaxIdleConns    int
	SetMaxOpenConns    int
	SetConnMaxLifetime time.Duration
}
