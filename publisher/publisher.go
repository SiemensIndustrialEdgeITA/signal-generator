package publisher

import (
	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
)

// Publisher streams data to the outside world
type Publisher interface {
	Start()
	Stop()
	Connect()
	SetIn(chan types.DataPoint)
	GetIn() chan types.DataPoint
}

func NewSimplePublisher(spc SimpleConfig) *SimpleSink {

	// Create mqtt client
	mclient := newMqttClient(spc.Mqtt)

	return &SimpleSink{
		cfg:  spc,
		Sink: mclient, // Inject the mqtt client
	}
}
