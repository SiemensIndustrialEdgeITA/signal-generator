package publisher

import (
	"fmt"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
	//	logger "github.com/sirupsen/logrus"
)

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

// Publisher stage factory
func NewPublisher(ttype pubType, c Config, mc MqttConfig) (Publisher, error) {
	q := make(chan struct{}) // Unbuffered

	switch ttype {
	case SIMPLE:
		var nc SimpleConfig = c.(SimpleConfig) // Assert config interface to Type
		return &SimpleSink{
			value:  0,
			coeff:  nc.Coeff,
			minVal: nc.MinVal,
			maxVal: nc.MaxVal,
			quit:   q,
		}, nil

	}
	return nil, fmt.Errorf("could not create publisher with type: %d", ttype)
}
