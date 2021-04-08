package generator

import (
	"time"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
	logger "github.com/sirupsen/logrus"
)

type SineConfig struct {
	SampleRate int     `c2s:"rate_ms"`
	Period     float64 `c2s:"period"`
	Amplitude  float64 `c2s:"amplitude"`
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
