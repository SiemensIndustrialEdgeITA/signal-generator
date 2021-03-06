package publisher

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	logger "github.com/sirupsen/logrus"
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
	jsonpayload, err := json.Marshal(msg.Payload)
	if err != nil {
		logger.Error("could not publish message: err:", err)
	}

	token := mc.client.Publish(msg.Topic, msg.Qos, msg.Retained, jsonpayload)
	token.Wait()
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	logger.Info("mqtt client connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	logger.Error("connect lost:", err.Error())
}
