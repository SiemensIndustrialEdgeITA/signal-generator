package publisher

import (
	"fmt"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
	//	logger "github.com/sirupsen/logrus"
)

// Publisher streams data to the outside world
type Publisher interface {
	Start()
	Stop()
	SetIn(chan types.DataPoint)
	GetIn() chan types.DataPoint
}

type Config interface{}

type pubType int

const (
	SIMPLE pubType = iota
)

// NewPublisher is Publisher stage factory
func NewPublisher(ttype pubType, c Config) (Publisher, error) {

	switch ttype {
	case SIMPLE:
		var sc SimpleConfig = c.(SimpleConfig) // Assert config interface to Type

		// Create mqtt client
		mclient := newMqttClient(sc.Mqtt)

		return &SimpleSink{
			cfg:  sc,
			Sink: mclient, // Inject the mqtt client
		}, nil
	}
	return nil, fmt.Errorf("could not create publisher with type: %d", ttype)
}
