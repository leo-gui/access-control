package internal

import (
	"access-control/config"
	"fmt"
	"github.com/spf13/viper"
)

func InitConfigByViper() config.Mysql{

	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config/config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	var _config *config.Server
	err = viper.Unmarshal(&_config)
	if err != nil {
		fmt.Println(err.Error())
	}

	return _config.Mysql

}

func InitConsByViper() config.Const{

	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config/config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	var _const *config.Server
	err = viper.Unmarshal(&_const)
	if err != nil {
		fmt.Println(err.Error())
	}

	return _const.Const
}

func InitRedisByViper() config.Redis{

	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config/config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	var _const *config.Server
	err = viper.Unmarshal(&_const)
	if err != nil {
		fmt.Println(err.Error())
	}

	return _const.Redis
}