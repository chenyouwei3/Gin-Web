package config

import (
	"github.com/spf13/viper"
)

func InitConfig(path, fileName, fileType string) error {
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(path) //开发环境
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&SystemConfig)
	if err != nil {
		return err
	}
	return nil
}

var SystemConfig = Config{}
