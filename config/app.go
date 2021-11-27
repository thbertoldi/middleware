package config

import (
	"github.com/spf13/viper"
)

// Influx holds all InfluxDB Configuration
type Influx struct {
	IFUser   string `mapstructure:"INFLUXDB_ADMIN_USER"`
	IFPW     string `mapstructure:"INFLUXDB_ADMIN_PASSWORD"`
	IFOrg    string `mapstructure:"INFLUXDB_ORG"`
	IFBucket string `mapstructure:"INFLUXDB_BUCKET"`
	IFToken  string `mapstructure:"TOKEN"`
	IFUrl    string `mapstructure:"URL"`
}

// MQTT Holds MQTT Configuration
type MQTT struct {
	Host  string   `json:"host"`
	Port  int      `json:"port"`
	Topic []string `json:"topic"`
	ID    string   `json:"id"`
}

// HTTP Holds MQTT Configuration
type HTTP struct {
	Port int `json:"port"`
}

// Internal holds all Application Configuration
type Internal struct {
	MQTT MQTT `json:"mqtt"`
	HTTP HTTP `json:"http"`
}

// LoadInternal loads internal configuration
func LoadInternal() (config Internal, err error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("app")
	v.SetConfigType("json")

	v.AutomaticEnv()

	err = v.ReadInConfig()
	if err != nil {
		return
	}

	err = v.Unmarshal(&config)
	return
}

// LoadInflux returns a structure with viper config
func LoadInflux() (config Influx, err error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")

	v.AutomaticEnv()

	err = v.ReadInConfig()
	if err != nil {
		return
	}

	err = v.Unmarshal(&config)
	return
}
