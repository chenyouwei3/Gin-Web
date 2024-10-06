package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/initialize/config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("initConfig err:", err)
		return
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}

var Conf Config

type Config struct {
	APP   APP   `yaml:"app"`
	MySQL MySQL `yaml:"mysql"`
}

type APP struct {
	Name string
	IP   string
	Port int
	Mode string
}

type MySQL struct {
	DriverName string `yaml:"driverName"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Database   string `yaml:"database"`
	UserName   string `yaml:"username"`
	Password   string `yaml:"password"`
	Charset    string `yaml:"charset"`
}
