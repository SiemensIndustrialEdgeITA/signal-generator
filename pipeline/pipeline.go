package pipeline

import (
	"fmt"
	"time"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/generator"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/transform"
	//	logger "github.com/sirupsen/logrus"
)

type Pipeline interface {
	Start()
	Stop()
	AddGenerator(generator.Generator)
	AddTransform(transform.Transform)
}

type datapoint struct {
	key string
	ts  time.Time
	val float64
}

func NewPipeline() (Pipeline, error) {
	q := make(chan struct{}) // Unbuffered

	return &DataPipeline{
		value:  0,
		coeff:  1,
		minVal: c.MinVal,
		maxVal: c.MaxVal,
		quit:   q,
	}, nil

}

type Coupling struct {
	In  chan datapoint
	Out chan datapoint
}

type DataPipeline struct {
	log *logger.Logger
}

func (n *NoiseTransform) Start() {
	fmt.Println("starting noise transform")

	for {
		select {
		case msg := <-n.In:
			fmt.Println("key:", msg.key)
			fmt.Println("ts:", msg.ts)
			fmt.Println("value:", msg.val)
			msg.val = 99
			n.Out <- msg
		case <-n.quit:
			fmt.Println("noisetransform: received close")
			return
		}
	}
}

func (n *NoiseTransform) Stop() {
	close(n.quit)
}

func (n *NoiseTransform) WireInput(in chan datapoint) {
	n.In = in
}

func (n *NoiseTransform) WireOutput(out chan datapoint) {
	n.Out = out
}
