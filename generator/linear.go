package generator

import (
	"time"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
	logger "github.com/sirupsen/logrus"
)

type LinearConfig struct {
	SampleRate time.Duration `mapstructure:"rate_ms"`
	Coeff      float64       `mapstructure:"coeff"`
	MinVal     float64       `mapstructure:"min"`
	MaxVal     float64       `mapstructure:"max"`
}

type LinearGenerator struct {
	cfg    LinearConfig
	value  float64
	ticker *time.Ticker
	quit   chan struct{}
	Out    chan types.DataPoint
}

func (l *LinearGenerator) Start() {
	logger.Info("starting linear generation")
	for {
		select {
		case t := <-l.ticker.C:
			intervalsec := float64(l.cfg.SampleRate.Seconds())
			l.value = l.value + (l.cfg.Coeff * intervalsec)
			if l.value > l.cfg.MaxVal {
				l.value = l.cfg.MinVal
			}
			msg := types.DataPoint{
				Key: "linear",
				Ts:  t,
				Val: l.value,
			}
			if len(l.Out) < (cap(l.Out) - 1) {
				l.Out <- msg
				logger.Info("generator: msg : { Key:", msg.Key, ", Ts:", msg.Ts, ", Val:", msg.Val, " }")
			} else {
				logger.Error("generator: out: max capacity exceeded: cap:", cap(l.Out), " len:", len(l.Out))
			}
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
