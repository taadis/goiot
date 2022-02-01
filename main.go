package main

import (
	"fmt"
	"log"

	"github.com/taadis/goiot/handler"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	var host = "broker.emqx.io"
	var port = 1883
	server := fmt.Sprintf("tcp://%s:%d", host, port)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(server)
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("emqx")
	opts.SetPassword("public")
	opts.SetDefaultPublishHandler(handler.MessagePublishHandler)
	opts.OnConnect = handler.OnConnect
	opts.OnConnectionLost = handler.OnConnectionLost
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	handler.Sub(client)
	handler.Publish(client)

	client.Disconnect(1000)
}
