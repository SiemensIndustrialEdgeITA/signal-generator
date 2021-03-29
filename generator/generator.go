package generator

import (
	"time"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
)

type Generator interface {
	Start()
	Stop()
	SetOut(chan types.DataPoint)
	GetOut() chan types.DataPoint
}

func NewLinearGen(lgc LinearConfig) *LinearGenerator {
	q := make(chan struct{}) // Unbuffered quit channel
	t := time.NewTicker(time.Duration(lgc.SampleRate) * time.Millisecond)
	return &LinearGenerator{
		cfg:    lgc,
		value:  0,
		ticker: t,
		quit:   q,
	}
}

func NewSineGen(sgc SineConfig) *SineGenerator {
	q := make(chan struct{}) // Unbuffered quit channel
	t := time.NewTicker(time.Duration(sgc.SampleRate) * time.Millisecond)
	return &SineGenerator{
		cfg:    sgc,
		value:  0,
		ticker: t,
		quit:   q,
	}
}
