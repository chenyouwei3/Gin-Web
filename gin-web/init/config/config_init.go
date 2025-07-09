package config

import (
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../config") //开发环境
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		return err
	}
	return nil
}

var Conf = Config{}
