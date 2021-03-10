package generator

import (
	"time"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
	logger "github.com/sirupsen/logrus"
)

type LinearConfig struct {
	SampleRate time.Duration
	Coeff      float64
	MinVal     int
	MaxVal     int
}

type LinearGenerator struct {
	log    *logger.Logger
	cfg    LinearConfig
	value  float64
	ticker *time.Ticker
	quit   chan struct{}
	Out    chan types.DataPoint
}

func (l *LinearGenerator) Start() {
	logger.Info("starting linear generation")
	logger.Info("interval:", l.cfg.SampleRate.Milliseconds())
	for {
		select {
		case t := <-l.ticker.C:
			intervalsec := float64(l.cfg.SampleRate.Seconds())
			l.value = l.value + (l.cfg.Coeff * intervalsec)
			msg := types.DataPoint{
				Key: "linear",
				Ts:  t,
				Val: l.value,
			}
			l.Out <- msg
		case <-l.quit:
			logger.Info("received close")
			l.ticker.Stop()
			return
		}
	}
}

func (l *LinearGenerator) Stop() {
	close(l.quit)
}

func (l *LinearGenerator) SetOut(o chan types.DataPoint) {
	l.Out = o
}

func (l *LinearGenerator) GetOut() chan types.DataPoint {
	return l.Out
}
