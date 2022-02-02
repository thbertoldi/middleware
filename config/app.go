package config

import (
	"bytes"
	"fmt"

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

var v1 = viper.New()

func Write(config []byte) error {
	fmt.Println(string(config))
	err := v1.ReadConfig(bytes.NewBuffer(config))
	if err != nil {
		return err
	}
	fmt.Println(v1.AllSettings())
	return v1.WriteConfig()
}

// LoadInternal loads internal configuration
func LoadInternal() (config Internal, err error) {
	v1.AddConfigPath(".")
	v1.SetConfigName("app")
	v1.SetConfigType("json")

	v1.AutomaticEnv()

	err = v1.ReadInConfig()
	if err != nil {
		return
	}

	err = v1.Unmarshal(&config)
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
