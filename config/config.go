package config

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	DefaultOptionConfig 	= 	&DefaultConfig{}
	DefaultConfigPath 		string
	DefaultConfigName 		string
	DefaultConfigType 		string

	once sync.Once
)

type DefaultConfig struct {

}

func InitConfig() *DefaultConfig{
	once.Do(func() {
		LoadConfig()
	})
	return DefaultOptionConfig
}

func LoadConfig(){
	viper.AddConfigPath(DefaultConfigPath)
	viper.SetConfigName(DefaultConfigName)
	viper.SetConfigType(DefaultConfigType)

	err := viper.ReadInConfig()
	if err != nil {
		panic("LoadConfigReadInConfigError:" + err.Error())
	}

	err = viper.Unmarshal(&DefaultOptionConfig)
	if err != nil {
		panic("LoadConfigUnmarshalError:"+err.Error())
	}
}