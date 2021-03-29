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

// BuildGenFromMap builds the specific generator type
func (ppl *Pipeline) BuildGenFromMap(gencfgmap StageCfgMap) (generator.Generator, error) {

	var gen generator.Generator

	switch gencfgmap.Type {
	case "linear":
		{
			gencfg, err := ParseLinGenCfg(gencfgmap.RawConf)
			if err != nil {
				return nil, fmt.Errorf("build generator: %s", err)
			}
			gen = generator.NewLinearGen(*gencfg)
		}
	default:
		{
			return nil, fmt.Errorf("build generator: could not find type %s", gencfgmap.Type)
		}
	}

	return gen, nil
}

// BuildTransFromMap builds the specific transform type
func (ppl *Pipeline) BuildTransFromMap(transcfgmap StageCfgMap) (transform.Transform, error) {

	var trans transform.Transform

	switch transcfgmap.Type {
	case "noise":
		{
			transcfg, err := ParseNoiseTransCfg(transcfgmap.RawConf)
			if err != nil {
				return nil, fmt.Errorf("build generator: %s", err)
			}
			trans = transform.NewNoiseTrans(*transcfg)
		}
	default:
		{
			return nil, fmt.Errorf("build transform: could not find type %s", transcfgmap.Type)
		}
	}

	return trans, nil
}

// BuildPubFromMap builds the specific publisher type
func (ppl *Pipeline) BuildPubFromMap(pubcfgmap StageCfgMap) (publisher.Publisher, error) {

	var pub publisher.Publisher
	mqttcfg := publisher.MqttConfig{
		Host:     "ie-databus",
		Port:     1883,
		User:     "simatic",
		Password: "simatic",
		ClientId: "signal-generator",
	}

	switch pubcfgmap.Type {
	case "simple":
		{
			pubcfg, err := ParseSimplePubCfg(pubcfgmap.RawConf)
			if err != nil {
				return nil, fmt.Errorf("build generator: %s", err)
			}
			pubcfg.Mqtt = mqttcfg
			pub = publisher.NewSimplePublisher(*pubcfg)
		}
	default:
		{
			return nil, fmt.Errorf("build publisher: could not find type %s", pubcfgmap.Type)
		}
	}

	return pub, nil
}

func (ppl *Pipeline) AddGenerator(gen *generator.Generator) {
	ppl.Gen = *gen
}

func (ppl *Pipeline) AddTransform(trans *transform.Transform) {
	ppl.Gen = *trans
}

func (ppl *Pipeline) AddPublisher(pub *publisher.Publisher) {
	ppl.Pub = *pub
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

	// Start publisher in parallel goroutine
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
