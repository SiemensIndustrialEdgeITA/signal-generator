package pipelines

import (
	"fmt"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/generator"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/publisher"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/transform"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/types"
	logger "github.com/sirupsen/logrus"
)

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
func (ppl *Pipeline) BuildGenerator(gencfg DynGenConfig) (generator.Generator, error) {

	var gen generator.Generator

	switch gencfg.Type {
	case "linear":
		{
			gencfg, ok := gencfg.RawConf.(generator.LinearConfig)
			if !ok {
				return nil, fmt.Errorf("cannot assert linear")
			}
			gen = generator.NewLinearGen(gencfg)
		}
	default:
		{
			return nil, fmt.Errorf("could not find type %s", gencfg.Type)
		}
	}

	return gen, nil
}

// BuildTransFromMap builds the specific transform type
func (ppl *Pipeline) BuildTransform(transcfg DynTransConfig) (transform.Transform, error) {

	var trans transform.Transform

	switch transcfg.Type {
	case "noise":
		{
			transcfg, ok := transcfg.RawConf.(transform.NoiseConfig)
			if !ok {
				return nil, fmt.Errorf("cannot assert noise")
			}
			trans = transform.NewNoiseTrans(transcfg)
		}
	default:
		{
			return nil, fmt.Errorf("could not find type %s", transcfg.Type)
		}
	}

	return trans, nil
}

// BuildPubFromMap builds the specific publisher type
func (ppl *Pipeline) BuildPublisher(sinkcfg DynSinkConfig) (publisher.Publisher, error) {

	var pub publisher.Publisher
	mqttcfg := publisher.MqttConfig{
		Host:     "127.0.0.1",
		Port:     1883,
		User:     "simatic",
		Password: "simatic",
		ClientId: "signal-generator-" + ppl.cfg.Name,
	}

	switch sinkcfg.Type {
	case "simple":
		{
			// Simple config is not read from file
			pubcfg := publisher.SimpleConfig{}
			pubcfg.Mqtt = mqttcfg
			pub = publisher.NewSimplePublisher(pubcfg)
		}
	default:
		{
			return nil, fmt.Errorf("could not find type %s", sinkcfg.Type)
		}
	}

	return pub, nil
}

func (ppl *Pipeline) AddGenerator(gen *generator.Generator) {
	logger.Info("add generator pipeline: ", ppl.cfg.Name)
	ppl.Gen = *gen
}

func (ppl *Pipeline) AddTransform(trans *transform.Transform) {
	logger.Info("add transform pipeline: ", ppl.cfg.Name)
	ppl.Trans = *trans
}

func (ppl *Pipeline) AddPublisher(pub *publisher.Publisher) {
	logger.Info("add publisher pipeline: ", ppl.cfg.Name)
	ppl.Pub = *pub
}

// Connect connects the whole pipeline stages
func (ppl *Pipeline) Connect() {

	logger.Info("connecting pipeline: ", ppl.cfg.Name)

	// Wire up stages with channnels
	// gen -> c1 -> tr -> c2 -> pub
	ppl.Gen.SetOut(ppl.GenOut)
	ppl.Trans.SetIn(ppl.GenOut)
	ppl.Trans.SetOut(ppl.TransOut)
	ppl.Pub.SetIn(ppl.TransOut)
	ppl.Pub.Connect()
}

func (ppl *Pipeline) Start() {
	logger.Info("starting pipeline: ", ppl.cfg.Name)

	// Start publisher in parallel goroutine
	go ppl.Pub.Start()

	// Start noise transform in parallel goroutine
	go ppl.Trans.Start()

	// Start data generation in parallel goroutine
	go ppl.Gen.Start()

}

func (ppl *Pipeline) Stop() {
	logger.Info("stopping pipeline: ", ppl.cfg.Name)

	ppl.Gen.Stop()
	ppl.Trans.Stop()
	ppl.Pub.Stop()
}
