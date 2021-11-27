package internal

import (
	"github.com/spf13/viper"
)

// Config holds all Application Configuration
type Config struct {
	Host       string `mapstructure:"MQ_HOST"`
	Port       int    `mapstructure:"MQ_PORT"`
	Topic      string `mapstructure:"MQ_TOPIC"`
	ID         string `mapstructure:"MQ_ID"`
	ServerPort int    `mapstructure:"SERVER_PORT"`
	IFUser     string `mapstructure:"INFLUXDB_ADMIN_USER"`
	IFPW       string `mapstructure:"INFLUXDB_ADMIN_PASSWORD"`
	IFOrg      string `mapstructure:"INFLUXDB_ORG"`
	IFBucket   string `mapstructure:"INFLUXDB_BUCKET"`
	IFToken    string `mapstructure:"TOKEN"`
	IFUrl      string `mapstructure:"URL"`
}

// LoadConfig returns a structure with viper config
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
