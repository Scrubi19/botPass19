package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	// Database config
	DbUsername             string `mapstructure;"db_username"`
	DbPassword             string `mapstructure;"db_password"`
	DbHost                 string `mapstructure;"db_host"`
	DbPort                 string `mapstructure;"db_port"`
	DbName                 string `mapstructure;"db_name"`
	DbSchema               string `mapstructure;"db_schema"`
	DbPreferSimpleProtocol bool   `mapstructure;"db_prefer_simple_protocol"`
	DbPrepareStatement     bool   `mapstructure;"db_prepare_statement"`
	DbMigrate              bool   `mapstructure;"db_migrate"`

	// Telegram
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
