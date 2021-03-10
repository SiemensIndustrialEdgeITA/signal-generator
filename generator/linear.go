package generator

import (
	"fmt"
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
	fmt.Println("starting linear generation")
	fmt.Println("interval:", l.cfg.SampleRate.Milliseconds())
	for {
		select {
		case t := <-l.ticker.C:
			//			fmt.Println("key:", "linear")
			//			fmt.Println("ts:", t)
			//			fmt.Println("value:", l.value)
			intervalsec := float64(l.cfg.SampleRate.Seconds())
			l.value = l.value + (l.cfg.Coeff * intervalsec)
			msg := types.DataPoint{
				Key: "linear",
				Ts:  t,
				Val: l.value,
			}
			l.Out <- msg
		case <-l.quit:
			fmt.Println("received close")
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
