package pipeline

import (
	logger "github.com/sirupsen/logrus"
	"os"
)

type PipesArray struct {
	Pipes []Pipeline
}

func NewPipeArray(cfgmap interface{}) (*PipesArray, error) {

	parr := &PipesArray{}

	cfg, err := ParseConfig(cfgmap)
	if err != nil {
		logger.Error("pipearray: %s", err)
		os.Exit(1)
	}

	for _, pipecfg := range cfg.PipesCfgMap {

		pipe := NewPipeline(PipeConfig{Name: pipecfg.NameCfgMap})

		// Build and assign the generator
		gen, err := pipe.BuildGenFromMap(pipecfg.GenCfgMap)
		if err != nil {
			logger.Error("pipeline %s: %s", pipe.cfg.Name, err)
			os.Exit(1)
		}
		pipe.AddGenerator(gen)

		// Build and assign the transform
		trans, err := pipe.BuildTransFromMap(pipecfg.TransCfgMap)
		if err != nil {
			logger.Error("pipeline %s: %s", pipe.cfg.Name, err)
			os.Exit(1)
		}
		pipe.AddTransform(trans)

		// Build and assign the publisher
		pub, err := pipe.BuildPubFromMap(pipecfg.GenCfgMap)
		if err != nil {
			logger.Error("pipeline %s: %s", pipe.cfg.Name, err)
			os.Exit(1)
		}
		pipe.AddPublisher(pub)

		// Connect the stages
		pipe.Connect()

		parr.Pipes = append(parr.Pipes, pipe)
	}

	return parr, nil
}

func (parr *PipesArray) Start() {
	logger.Info("Starting pipelines array")

	for _, pipe := range parr.Pipes {
		pipe.Start()
	}
}

func (parr *PipesArray) Stop() {
	logger.Info("Stopping pipelines array")

	for _, pipe := range parr.Pipes {
		pipe.Stop()
	}
}
