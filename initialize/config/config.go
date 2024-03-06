package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

// 系统配置，对应yml
// 声明全局变量
var Config = Init("././test.yaml")

type config struct {
	Mysql   *MysqlConfig   `yaml:"mysql"`
	Mongodb *MongodbConfig `yaml:"mongodb"`
	Redis   *RedisConfig   `yaml:"redis"`
	Jwt     *JwtConfig     `yaml:"jwt"`
}

type MysqlConfig struct {
	Address            string        `yaml:"Address"`            //地址
	SetMaxIdleConns    int           `yaml:"SetMaxIdleConns"`    //最大空闲连接数
	SetMaxOpenConns    int           `yaml:"SetMaxOpenConns"`    //最大连接数
	SetConnMaxLifetime time.Duration `yaml:"SetConnMaxLifetime"` //设置连接空闲超时
}

type MongodbConfig struct {
	Address   string `yaml:"Address"`   //地址
	DBName001 string `yaml:"DBName001"` //数据库名字
	DBName002 string `yaml:"DBName002"` //数据库名字
}

type RedisConfig struct {
	Address      string `yaml:"Address"`      // Redis 服务器地址
	Password     string `yaml:"Password"`     // Redis 服务器密码
	DB           int    `yaml:"DB"`           // Redis 数据库索引
	PoolSize     int    `yaml:"PoolSize"`     // 连接池大小
	MinIdleConns int    `yaml:"MinIdleConns"` // 最小空闲连接数
}

type JwtConfig struct {
	SignKey       string `yaml:"SignKey"`
	ExpireSeconds int    `yaml:"ExpireSeconds"`
	Issuer        string `yaml:"Issuer"`
}

func Init(filename string) *config {
	Config := &config{}
	if yamlFile, err := os.ReadFile(filename); err != nil {
		logrus.Fatalln("读取配置文件错误", err)
	} else if err = yaml.Unmarshal(yamlFile, Config); err != nil {
		logrus.Fatalln("读取配置文件错误", err)
	}
	return Config
}
