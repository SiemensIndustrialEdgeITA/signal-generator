package generator

import (
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
	cfg    SineConfig
	value  float64
	ticker *time.Ticker
	quit   chan struct{}
	Out    chan types.DataPoint
}

func (s *SineGenerator) Start() {
	logger.Info("starting sine generation")
	logger.Info("interval:", s.cfg.SampleRate)
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
