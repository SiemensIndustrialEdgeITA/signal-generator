package transform

import (
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

func NewNoiseTrans(ntc NoiseConfig) *NoiseTransform {
	q := make(chan struct{}) // Unbuffered quit channel
	return &NoiseTransform{
		cfg:   ntc,
		value: 0,
		quit:  q,
	}
}
