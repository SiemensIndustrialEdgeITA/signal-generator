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
	log      *logger.Logger
	interval time.Duration
	value    float64
	coeff    float64
	minVal   int
	maxVal   int
	ticker   *time.Ticker
	quit     chan struct{}
	Out      chan types.DataPoint
}

func (l *LinearGenerator) Start() {
	fmt.Println("starting linear generation")
	fmt.Println("interval:", l.interval)
	for {
		select {
		case t := <-l.ticker.C:
			//			fmt.Println("key:", "linear")
			//			fmt.Println("ts:", t)
			//			fmt.Println("value:", l.value)
			intervalsec := float64(l.interval) / 1000000000
			l.value = l.value + (l.coeff * intervalsec)
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
