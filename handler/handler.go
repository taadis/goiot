package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/taadis/goiot/model"
)

const topic string = "topic/test"
const topicGoiot string = "topic/goiot"

func OnConnect(client mqtt.Client) {
	log.Printf("Connected\n")
}

func OnConnectionLost(client mqtt.Client, err error) {
	log.Printf("Connection lost: %+v\n", err)
}

func MessagePublishHandler(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Publishing message: %s from topic:%s", msg.Payload(), msg.Topic())
}

// 消息订阅回调
func MessageSubscribeCallback(client mqtt.Client, message mqtt.Message) {
	var data model.Model
	err := json.Unmarshal(message.Payload(), &data)
	if err != nil {
		log.Printf("Subscribed to meesage json.Unmarshal error:%+v", err)
		return
	}

	log.Printf("Subscribed to message id:%d and payload:%+v", message.MessageID(), data)
}

func Sub(client mqtt.Client) {
	token := client.Subscribe(topicGoiot, 1, MessageSubscribeCallback)
	token.Wait()
	log.Printf("Subscribed to topic %s", topic)
}

func Publish(client mqtt.Client) {
	num := 10
	// for i := 0; i < num; i++ {
	// 	text := fmt.Sprintf("Message %d", i)
	// 	token := client.Publish(topic, 1, false, text)
	// 	token.Wait()
	// 	time.Sleep(time.Second)
	// }
	for i := 0; i < num; i++ {
		data := &model.Model{}
		data.IDnum = int64(i)
		data.MachineIP = fmt.Sprintf("127.0.0.%d", i)
		payload, err := json.Marshal(data)
		if err != nil {
			log.Printf("publish data json.Marshal error:%+v", err)
			continue
		}
		token := client.Publish(topicGoiot, 1, false, payload)
		token.Wait()
		time.Sleep(time.Second)
	}
}
