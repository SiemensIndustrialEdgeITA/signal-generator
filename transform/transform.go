package transform

import (
	"fmt"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
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

type TransType int

const (
	NOISE TransType = iota
)

// Transform stage factory
func NewTransform(ttype TransType, c Config) (Transform, error) {
	q := make(chan struct{}) // Unbuffered

	switch ttype {
	case NOISE:
		var nc NoiseConfig = c.(NoiseConfig) // Assert config interface to Type
		return &NoiseTransform{
			cfg:   nc,
			value: 0,
			quit:  q,
		}, nil

	}
	return nil, fmt.Errorf("could not create transform type: %d", ttype)
}
