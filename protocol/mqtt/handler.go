package mqtt

import (
	"fmt"

	paho "github.com/eclipse/paho.mqtt.golang"
)

var msgPubHandler paho.MessageHandler = func(client paho.Client, msg paho.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var onSensorData paho.MessageHandler = func(client paho.Client, msg paho.Message) {
	fmt.Printf("Received message from sensor: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connHandler paho.OnConnectHandler = func(client paho.Client) {
	fmt.Println("Connected")
}

var connLostHandler paho.ConnectionLostHandler = func(client paho.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}
