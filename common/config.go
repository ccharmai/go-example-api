package common

import (
	"log"

	"github.com/spf13/viper"
)

var Config ConfigModel

type ConfigModel struct {
	AccessTokenPrivateKey  string `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	Port                   string `mapstructure:"PORT"`
}

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Config file fatal error: ", err.Error())
	}

	viper.Unmarshal(&Config)
}
