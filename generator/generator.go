package generator

import (
	"fmt"
	"time"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
	//	logger "github.com/sirupsen/logrus"
)

type Generator interface {
	Start()
	Stop()
	SetOut(chan types.DataPoint)
	GetOut() chan types.DataPoint
}

type Config interface{}

type gentype int

const (
	LINEAR gentype = iota
	SINE
)

// Generator factory
func NewGenerator(gtype gentype, c Config) (Generator, error) {
	q := make(chan struct{}) // Unbuffered

	switch gtype {
	case LINEAR:
		var lc LinearConfig = c.(LinearConfig) // Assert config interface to Type
		t := time.NewTicker(lc.SampleRate)
		return &LinearGenerator{
			cfg:    lc,
			value:  0,
			ticker: t,
			quit:   q,
		}, nil
	case SINE:
		var sc SineConfig = c.(SineConfig) // Assert config interface to Type
		t := time.NewTicker(sc.SampleRate)
		return &SineGenerator{
			cfg:    sc,
			value:  0,
			ticker: t,
			quit:   q,
		}, nil

	}
	return nil, fmt.Errorf("could not create generator type: %d", gtype)
}
