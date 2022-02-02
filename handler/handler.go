package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	influxdb "github.com/influxdata/influxdb-client-go/v2"
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

	// write to influxdb
	writeToInflux(&data)
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
		data.MachineSN = fmt.Sprintf("SN%d", i)
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

func connInflux() influxdb.Client {
	server := "http://localhost:8086"
	authToken := "dev-token"
	client := influxdb.NewClient(server, authToken)
	_, err := client.Ping(context.Background())
	if err != nil {
		log.Printf("influxdb ping error:%+v", err)
	}

	return client
}

func writeToInflux(data *model.Model) error {
	client := connInflux()
	defer client.Close()

	org := "taadis"
	bucket := "my-bucket"
	writeAPI := client.WriteAPIBlocking(org, bucket)

	point := influxdb.NewPoint("machine_info",
		map[string]string{"machine_sn": data.MachineSN},
		map[string]interface{}{"id_num": data.IDnum, "machine_ip": data.MachineIP},
		time.Now(),
	)
	// write point immediately
	err := writeAPI.WritePoint(context.Background(), point)
	if err != nil {
		log.Printf("influxdb write point error:%+v", err)
		return err
	}
	log.Printf("influxdb write point success")
	return nil
}
