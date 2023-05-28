package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	BotApiToken string `mapstructure:"bot_api_token"`
}

func Init() (config *Config, err error) {
	viper.AddConfigPath("config/")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	err = viper.Unmarshal(&config)
	return
}
