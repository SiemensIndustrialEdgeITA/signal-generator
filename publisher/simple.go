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

func (s *SimpleSink) Start() {
	logger.Info("starting simplesink publisher")

	s.Sink.Connect()

	for {
		select {
		case inmsg := <-s.In:
			outmsg := types.DataPoint{
				Key: inmsg.Key,
				Ts:  inmsg.Ts,
				Val: inmsg.Val,
			}
			s.Sink.Publish(&MqttMsg{
				Topic:    "simple",
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
