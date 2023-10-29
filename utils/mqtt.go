package utils

import (
	"log"
	"os"

	paho "github.com/eclipse/paho.mqtt.golang"
)

var client paho.Client

func Init() {
	opts := paho.NewClientOptions()
	opts.AddBroker(os.Getenv("MQTT_BROKER")) // e.g. "tcp://broker.hivemq.com:1883"
	opts.SetClientID(os.Getenv("MQTT_CLIENT_ID"))

	client = paho.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to MQTT Broker: %v", token.Error())
	}
}

func Publish(topic string, payload string) {
	token := client.Publish(topic, 0, false, payload)
	token.Wait()
}
