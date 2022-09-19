package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB_SOURCE     string `mapstructure:"DB_SOURCE"`
	REDIS_SOURCE  string `mapstructure:"REDIS_SOURCE"`
	QN_ACCESS_KEY string `mapstructure:"QN_ACCESS_KEY"`
	QN_SECRET_KEY string `mapstructure:"QN_SECRET_KEY"`
	QN_BUCKET     string `mapstructure:"QN_BUCKET"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// read variable from env variable
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
