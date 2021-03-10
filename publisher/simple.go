package publisher

import (
	"fmt"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
	logger "github.com/sirupsen/logrus"
)

type SimpleConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

type SimpleSink struct {
	log    *logger.Logger
	Schema types.DataPoint
	quit   chan struct{}
	In     chan types.DataPoint
	Sink   MqttClient
}

func (s *SimpleSink) Start() {
	fmt.Println("starting simplesink publisher")
	for {
		select {
		case inmsg := <-s.In:
			outmsg := types.DataPoint{
				Key: inmsg.Key,
				Ts:  inmsg.Ts,
				Val: newVal,
			}
		case <-s.quit:
			fmt.Println("received close")
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
