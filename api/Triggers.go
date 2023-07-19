package api

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// mqtt trigger ------------------------------------------------------- MQTT
func connect(clientId string, settings MqttSettings) mqtt.Client {
	opts := createClientOptions(clientId, settings)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}
func createClientOptions(clientId string,settings MqttSettings) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s://%s:%s", settings.Protocol, settings.Host, settings.Port))
	opts.SetClientID(clientId)
	opts.SetUsername(settings.Username)
	opts.SetPassword(settings.Password)
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

func MqttTrigger(flow Flow, outputChannel chan interface{}) {
	mqttSettingsJson := flow.Triggers[0].Settings
	var mqttSettings MqttSettings
	errUnmarshal := json.Unmarshal([]byte(mqttSettingsJson), &mqttSettings)
	if errUnmarshal != nil {
	  fmt.Println(errUnmarshal)
	  return
	}

	client := connect("mqttTrigger", mqttSettings)
	client.Subscribe(mqttSettings.Topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		//fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
		outputChannel <- string(msg.Payload())
	})
}
