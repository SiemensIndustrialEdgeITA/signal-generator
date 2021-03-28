package pipeline

import (
	"fmt"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/generator"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/publisher"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/transform"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
	logger "github.com/sirupsen/logrus"
)

type PipeConfig struct {
	Name string
}

type Pipeline struct {
	cfg      PipeConfig
	Gen      generator.Generator
	Trans    transform.Transform
	Pub      publisher.Publisher
	GenOut   chan types.DataPoint
	TransOut chan types.DataPoint
}

func NewPipeline(pcfg PipeConfig) Pipeline {
	gout := make(chan types.DataPoint, 100) // Buffered
	tout := make(chan types.DataPoint, 100) // Buffered

	return Pipeline{
		cfg:      pcfg,
		GenOut:   gout,
		TransOut: tout,
	}
}

// BuildGenerator builds the specific generator type
func (ppl *Pipeline) BuildGenFromMap(gencfg StageCfgMap) (generator.Generator, error) {

	var gen generator.Generator

	switch gencfg.Type {
	case "linear":
		{
			gencfg, err := ParseLinGenCfg(gencfg.RawConf)
			if err != nil {
				return nil, fmt.Errorf("build generator: %s", err)
			}
			gen = generator.NewLinearGen(*gencfg)
		}
	default:
		{
			return nil, fmt.Errorf("build generator: could not find type %s", gencfg.Type)
		}
	}

	return gen, nil
}

func (ppl *Pipeline) AddGenerator(gen generator.Generator) {
	ppl.Gen = gen
}

func (ppl *Pipeline) AddTransform(trans transform.Transform) {
	ppl.Gen = trans
}

func (ppl *Pipeline) AddPublisher(pub publisher.Publisher) {
	ppl.Pub = pub
}

// Connect connects the whole pipeline stages
func (ppl *Pipeline) Connect() {

	// Wire up stages with channnels
	// gen -> c1 -> tr -> c2 -> pub
	ppl.Gen.SetOut(ppl.GenOut)
	ppl.Trans.SetIn(ppl.GenOut)
	ppl.Trans.SetOut(ppl.TransOut)
	ppl.Pub.SetIn(ppl.TransOut)
	ppl.Pub.Connect()
}

func (ppl *Pipeline) Start() {
	logger.Info("starting pipeline")

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
