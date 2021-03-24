package publisher

import (
	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
	logger "github.com/sirupsen/logrus"
)

type SimpleConfig struct {
	Mqtt MqttConfig
}

type SimpleSink struct {
	log  *logger.Logger
	cfg  SimpleConfig
	quit chan struct{}
	In   chan types.DataPoint
	Sink mqttClient
}

func (s *SimpleSink) Connect() {
	s.Sink.Connect()
}

func (s *SimpleSink) Start() {
	logger.Info("starting simplesink publisher")

	for {
		select {
		case inmsg := <-s.In:
			outmsg := types.DataPoint{
				Key: inmsg.Key,
				Ts:  inmsg.Ts,
				Val: inmsg.Val,
			}
			logger.Info("publisher: msg : { Key:", outmsg.Key, ", Ts:", outmsg.Ts, ", Val:", outmsg.Val, " }")
			s.Sink.Publish(&MqttMsg{
				Topic:    "signal-generator/simplejson",
				Qos:      0,
				Retained: false,
				Payload:  outmsg,
			})
		case <-s.quit:
			logger.Info("received close")
			return
		}
	}
}

func (s *SimpleSink) Stop() {
	close(s.quit)
}

func (s *SimpleSink) SetIn(in chan types.DataPoint) {
	s.In = in
}

func (s *SimpleSink) GetIn() chan types.DataPoint {
	return s.In
}
