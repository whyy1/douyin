package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB_SOURCE             string `mapstructure:"DB_SOURCE"`
	OSS_ACCESS_KEY_ID     string `mapstructure:"OSS_ACCESS_KEY_ID"`
	OSS_ACCESS_KEY_SECRET string `mapstructure:"OSS_ACCESS_KEY_SECRET"`
	OSS_ENDPOINT          string `mapstructure:"OSS_ENDPOINT"`
	OSS_BUCKET            string `mapstructure:"OSS_BUCKET"`
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
