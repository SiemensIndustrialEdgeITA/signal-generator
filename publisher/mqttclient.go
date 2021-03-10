package publisher

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttMsg struct {
	Topic    string
	Qos      byte
	Retained bool
	Payload  interface{}
}

type MqttConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	ClientId string
}

type mqttClient struct {
	cfg    MqttConfig
	client mqtt.Client
}

func newMqttClient(mqttconf MqttConfig) mqttClient {

	newclient := mqttClient{
		cfg: mqttconf,
	}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", mqttconf.Host, mqttconf.Port))
	opts.SetClientID(mqttconf.ClientId)
	opts.SetUsername(mqttconf.User)
	opts.SetPassword(mqttconf.Password)
	//	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	newclient.client = mqtt.NewClient(opts)
	return newclient
}

func (mc *mqttClient) Connect() {
	if token := mc.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (mc *mqttClient) Publish(msg *MqttMsg) {
	token := mc.client.Publish(msg.Topic, msg.Qos, msg.Retained, msg.Payload)
	token.Wait()
}

//var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
//	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
//}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}
