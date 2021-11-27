package influx

import (
	"context"
	"fmt"
	"time"

	"codigos.ufsc.br/g.manoel/pi_das_2021_2/config"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// Model represents an InfluxDB instance
type Model struct {
	cli influxdb2.Client
	wri api.WriteAPIBlocking
}

// Get returns a Influx Model
func Get() Model {
	conf, _ := config.LoadInflux()
	client := influxdb2.NewClient(conf.IFUrl, conf.IFToken)
	writeAPI := client.WriteAPIBlocking(conf.IFOrg, conf.IFBucket)
	return Model{client, writeAPI}

}

// PerformPost return influx client
func (s Model) PerformPost(device string, data map[string]interface{}) {
	p := influxdb2.NewPoint("LMM",
		map[string]string{"device": device},
		data,
		time.Now())
	err := s.wri.WritePoint(context.Background(), p)
	if err != nil {
		fmt.Println("Error on Point write")
	}
}
