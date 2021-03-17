package pipeline

import (
	"fmt"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/generator"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/publisher"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/transform"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
	//	logger "github.com/sirupsen/logrus"
)

type Pipe interface {
	Start()
	Stop()
	Build()
	AddGenerator(generator.Config) error
	AddTransform(transform.Config) error
	AddPublisher(publisher.Config) error
}

type Pipeline struct {
	Gen      generator.Generator
	Trans    transform.Transform
	Pub      publisher.Publisher
	GenOut   chan types.DataPoint
	TransOut chan types.DataPoint
}

func NewPipeline() (*Pipeline, error) {
	gout := make(chan types.DataPoint, 100) // Buffered
	tout := make(chan types.DataPoint, 100) // Buffered

	return &Pipeline{
		GenOut:   gout,
		TransOut: tout,
	}, nil

}

func (ppl *Pipeline) AddGenerator(gtype generator.Gentype, gconf generator.Config) error {

	// Instance new data generator
	gen, err := generator.NewGenerator(gtype, gconf)
	if err != nil {
		return fmt.Errorf("could not add the generator")
	}

	ppl.Gen = gen

	return nil
}

func (ppl *Pipeline) AddTransform(ttype transform.TransType, tconf transform.Config) error {

	// Instance new data transform
	trns, err := transform.NewTransform(ttype, tconf)
	if err != nil {
		return fmt.Errorf("could not add the transform")
	}

	ppl.Trans = trns

	return nil
}

func (ppl *Pipeline) AddPublisher(ptype publisher.PubType, pconf publisher.Config) error {

	// Instance new publisher
	pub, err := publisher.NewPublisher(ptype, pconf)
	if err != nil {
		return fmt.Errorf("could not add publisher")
	}

	ppl.Pub = pub

	return nil
}

func (ppl *Pipeline) Build() {

	// Wire up stages with channnels
	// gen -> c1 -> tr -> c2 -> pub
	ppl.Gen.SetOut(ppl.GenOut)
	ppl.Trans.SetIn(ppl.GenOut)
	ppl.Trans.SetOut(ppl.TransOut)
	ppl.Pub.SetIn(ppl.TransOut)
	ppl.Pub.Connect()
}

func (ppl *Pipeline) Start() {
	fmt.Println("starting pipeline")

	// Start publisher
	go ppl.Pub.Start()

	// Start noise transform in parallel goroutine
	go ppl.Trans.Start()

	// Start data generation in parallel goroutine
	go ppl.Gen.Start()

}

func (ppl *Pipeline) Stop() {
	ppl.Gen.Stop()
	ppl.Trans.Stop()
	ppl.Pub.Stop()
}
