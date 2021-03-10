package generator

import (
	"fmt"
	"time"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
	logger "github.com/sirupsen/logrus"
)

type SineConfig struct {
	SampleRate time.Duration
	Period     float64
	Amplitude  float64
}

type SineGenerator struct {
	log    *logger.Logger
	cfg    SineConfig
	value  float64
	ticker *time.Ticker
	quit   chan struct{}
	Out    chan types.DataPoint
}

func (s *SineGenerator) Start() {
	fmt.Println("starting sine generation")
	fmt.Println("interval:", s.cfg.SampleRate)
}

func (s *SineGenerator) Stop() {
	close(s.quit)
}

func (s *SineGenerator) SetOut(o chan types.DataPoint) {
	s.Out = o
}

func (s *SineGenerator) GetOut() chan types.DataPoint {
	return s.Out
}
