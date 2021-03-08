package transform

import (
	"fmt"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
	//	logger "github.com/sirupsen/logrus"
)

type NoiseConfig struct {
	Coeff  float64
	MinVal int
	MaxVal int
}

type NoiseTransform struct {
	value  float64
	coeff  float64
	minVal int
	maxVal int
	quit   chan struct{}
	In     chan types.DataPoint
	Out    chan types.DataPoint
}

func (n *NoiseTransform) Start() {
	fmt.Println("starting noise transform")
	for {
		select {
		case inmsg := <-n.In:
			outmsg := types.DataPoint{
				Key: inmsg.Key,
				Ts:  inmsg.Ts,
				Val: 999,
			}
			n.Out <- outmsg
		case <-n.quit:
			fmt.Println("received close")
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
