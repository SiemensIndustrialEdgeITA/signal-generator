package main

type PipeConfig struct {
	Pipelines []Pipe `yaml:"pipelines"`
}

type Pipe struct {
	Name      string      `yaml:"name"`
	Generator StageConfig `yaml:"generator"`
}

type StageConfig struct {
	Type    string      `yaml:"type"`
	RawConf interface{} `yaml:"config"`
}

type GenConfig struct {
	Rate_ms int `mapstructure:"rate_ms"`
	Min     int `mapstructure:"min"`
	Max     int `mapstructure:"max"`
}

//type GenConfig struct {
//	Rate_ms int     `mapstructure:"rate_ms"`
//	Coeff   float64 `mapstructure:"coeff"`
//	Min     int     `mapstructure:"min"`
//	Max     int     `mapstructure:"max"`
//}
