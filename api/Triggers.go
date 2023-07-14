package api

import (
	"fmt"
	"log"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// mqtt trigger ------------------------------------------------------- MQTT
func connect(clientId string, host string, port int) mqtt.Client {
	opts := createClientOptions(clientId, host, port)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}
func createClientOptions(clientId string, host string, port int) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	protocol := "ws"
	opts.AddBroker(fmt.Sprintf("%s://%s:%d", protocol, host, port))
	opts.SetClientID(clientId)
	opts.SetUsername("!@3%4*N]ZY@KfqSJ")
	opts.SetPassword("9w#v;7Ma?*:5]W!U")
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	return opts
}
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	// fmt.Println("MQTT Connected")
}
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	// fmt.Printf("Connect lost: %v", err)
}

func listen(host string, port int, topic string, outputChannel chan string) {
	client := connect("mqttTrigger", host, port)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
		outputChannel <- string(msg.Payload())
	})
}

func MqttTrigger(mqttSettings MqttSettings, outputChannel chan string){
	topic := mqttSettings.Topic
	host := mqttSettings.Host
	port, _ := strconv.Atoi(mqttSettings.Port)

	go listen(host, port, topic, outputChannel)
}