package internal

import (
	"github.com/spf13/viper"
)

// Config holds all Application Configuration
type Config struct {
	Host       string `mapstructure:"HOST"`
	Port       int    `mapstructure:"PORT"`
	Topic      string `mapstructure:"TOPIC"`
	ID         string `mapstructure:"ID"`
	ServerPort int    `mapstructure:"SERVER_PORT"`
}

// LoadConfig returns a structure with viper config
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
