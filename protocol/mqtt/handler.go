package mqtt

import (
	"encoding/json"
	"fmt"
	"strings"

	"codigos.ufsc.br/g.manoel/pi_das_2021_2/protocol/influx"
	paho "github.com/eclipse/paho.mqtt.golang"
)

var influxModel influx.Model = influx.Get()

var msgPubHandler paho.MessageHandler = func(client paho.Client, msg paho.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var onSensorData paho.MessageHandler = func(client paho.Client, msg paho.Message) {
	fmt.Printf("Received message from sensor: %s from topic: %s\n", msg.Payload(), msg.Topic())
	var data map[string]interface{}
	message := strings.ReplaceAll(string(msg.Payload()), "'", "\"")
	err := json.Unmarshal([]byte(message), &data)
	if err != nil {
		fmt.Println("Failed to parse")
		return
	}
	topic := msg.Topic()
	elem := strings.Split(topic, " ")
	elem = elem[:len(elem)-1]
	influxModel.PerformPost(strings.Join(elem, " "), data)
}

var connHandler paho.OnConnectHandler = func(client paho.Client) {
	fmt.Println("Connected")
}

var connLostHandler paho.ConnectionLostHandler = func(client paho.Client, err error) {
	fmt.Printf("Connect lost: %v\n", err)
}
