package transform

import (
	"fmt"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
	//	logger "github.com/sirupsen/logrus"
)

type Transform interface {
	Start()
	Stop()
	SetIn(chan types.DataPoint)
	SetOut(chan types.DataPoint)
	GetIn() chan types.DataPoint
	GetOut() chan types.DataPoint
}

type Config interface{}

type transType int

const (
	NOISE transType = iota
)

// Transform stage factory
func NewTransform(ttype transType, c Config) (Transform, error) {
	q := make(chan struct{}) // Unbuffered

	switch ttype {
	case NOISE:
		var nc NoiseConfig = c.(NoiseConfig) // Assert config interface to Type
		return &NoiseTransform{
			value:  0,
			coeff:  nc.Coeff,
			minVal: nc.MinVal,
			maxVal: nc.MaxVal,
			quit:   q,
		}, nil

	}
	return nil, fmt.Errorf("could not create transform type: %d", ttype)
}
