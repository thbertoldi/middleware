package mqtt

import (
	"fmt"

	"codigos.ufsc.br/g.manoel/pi_das_2021_2/internal"
	paho "github.com/eclipse/paho.mqtt.golang"
)

// Run starts MQTT client
func Run(conf *internal.Config) {
	opts := paho.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", conf.Host, conf.Port))
	opts.SetClientID(conf.ID)
	opts.SetDefaultPublishHandler(msgPubHandler)
	opts.OnConnect = connHandler
	opts.OnConnectionLost = connLostHandler
	client := paho.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	sub(client, conf.Topic)
}

func sub(client paho.Client, topic string) {
	token := client.Subscribe(topic, 1, onSensorData)
	token.Wait()
	other := []string{"ESP32 AC 000", "ESP32 TF 000", "ESP32 TF 001", "ESP32 TF 002"}
	for _, t := range other {
		token := client.Subscribe(t+" envia", 1, onSensorData)
		token.Wait()
	}
}
