package pipelines

import (
	logger "github.com/sirupsen/logrus"
	"os"
)

type PipesArray struct {
	Pipes []Pipeline
}

func NewPipeArray(cfg *Config) (*PipesArray, error) {

	parr := &PipesArray{}

	for _, pipecfg := range cfg.PipeArrCfg {

		pipe := NewPipeline(pipecfg)

		// Build and assign the generator
		gen, err := pipe.BuildGenerator(pipecfg.GenCfg)
		if err != nil {
			logger.Error("pipeline:", pipe.cfg.Name, " err:", err)
			os.Exit(1)
		}
		pipe.AddGenerator(&gen)

		// Build and assign the transform
		trans, err := pipe.BuildTransform(pipecfg.TransfCfg)
		if err != nil {
			logger.Error("pipeline:", pipe.cfg.Name, " err:", err)
			os.Exit(1)
		}
		pipe.AddTransform(&trans)

		// Build and assign the publisher
		pub, err := pipe.BuildPublisher(pipecfg.SinkCfg)
		if err != nil {
			logger.Error("pipeline:", pipe.cfg.Name, " err:", err)
			os.Exit(1)
		}
		pipe.AddPublisher(&pub)

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
