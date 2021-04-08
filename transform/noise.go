package transform

import (
	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
	logger "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type NoiseConfig struct {
	Coeff  float64 `c2s:"coeff"`
	MinVal float64 `c2s:"min"`
	MaxVal float64 `c2s:"max"`
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
			if len(n.Out) < (cap(n.Out) - 1) {
				n.Out <- outmsg
				logger.Info("transform: msg : { Key:", outmsg.Key, ", Ts:", outmsg.Ts, ", Val:", outmsg.Val, " }")
			} else {
				logger.Error("transform: out: max capacity exceeded: cap:", cap(n.Out), " len:", len(n.Out))
			}
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
