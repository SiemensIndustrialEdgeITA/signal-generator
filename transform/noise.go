package transform

import (
	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
	logger "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type NoiseConfig struct {
	Coeff  float64
	MinVal float64
	MaxVal float64
}

type NoiseTransform struct {
	cfg   NoiseConfig
	value float64
	quit  chan struct{}
	In    chan types.DataPoint
	Out   chan types.DataPoint
}

func (n *NoiseTransform) Start() {
	logger.Info("starting noise transform")
	for {
		select {
		case inmsg := <-n.In:
			newVal := inmsg.Val + n.getNoiseForRange(n.cfg.MinVal, n.cfg.MaxVal)
			outmsg := types.DataPoint{
				Key: inmsg.Key,
				Ts:  inmsg.Ts,
				Val: newVal,
			}
			n.Out <- outmsg
		case <-n.quit:
			logger.Info("received close")
			return
		}
	}
}

func (n *NoiseTransform) Stop() {
	close(n.quit)
}

func (n *NoiseTransform) SetIn(in chan types.DataPoint) {
	n.In = in
}

func (n *NoiseTransform) GetIn() chan types.DataPoint {
	return n.In
}

func (n *NoiseTransform) SetOut(o chan types.DataPoint) {
	n.Out = o
}

func (n *NoiseTransform) GetOut() chan types.DataPoint {
	return n.Out
}

func (n *NoiseTransform) getNoiseForRange(min float64, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	noise := rand.Float64() * (max - min)
	return noise
}
