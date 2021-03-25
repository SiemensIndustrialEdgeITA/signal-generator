package main

type PipeConfig struct {
	Pipelines []Pipe `mapstructure:"pipelines"`
}

type Pipe struct {
	Name      string      `mapstructure:"name"`
	Generator StageConfig `mapstructure:"generator"`
}

type StageConfig struct {
	Type    string      `mapstructure:"type"`
	RawConf interface{} `mapstructure:"config"`
}

type GenConfig struct {
	Rate_ms int     `mapstructure:"rate_ms"`
	Coeff   float64 `mapstructure:"coeff"`
	Min     int     `mapstructure:"min"`
	Max     int     `mapstructure:"max"`
}
