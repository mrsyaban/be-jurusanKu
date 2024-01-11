package config

import (
	"time"

	"github.com/spf13/viper"
)


type Config struct {
	DB_DRIVER string `mapstructure:"DB_DRIVER"`
	DB_URL string `mapstructure:"DB_URL"`
	SERVER_ADDR string	`mapstructure:"SERVER_ADDR"`
	JWT_KEY string `mapstructure:"JWT_KEY"`
	JWT_EXP time.Duration `mapstructure:"JWT_EXP"`
}

func Init(path string) (config Config, err error) {
	var runtime_viper = viper.New()
	runtime_viper.AddConfigPath(path)
	runtime_viper.SetConfigName(".env")
	runtime_viper.SetConfigType("env")

	runtime_viper.AutomaticEnv()
	err = runtime_viper.ReadInConfig()

	if err != nil {
		return
	}

	err = runtime_viper.Unmarshal(&config)
	return
}
