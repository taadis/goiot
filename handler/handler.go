package handler

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const topic string = "topic/test"

func OnConnect(client mqtt.Client) {
	log.Printf("Connected\n")
}

func OnConnectionLost(client mqtt.Client, err error) {
	log.Printf("Connection lost: %+v\n", err)
}

func MessagePublishHandler(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message: %s from topic:%s", msg.Payload(), msg.Topic())
}

func Sub(client mqtt.Client) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	log.Printf("Subscribed to topic %s", topic)
}

func Publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish(topic, 1, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}
