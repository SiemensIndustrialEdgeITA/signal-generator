package pipeline

import (
	"fmt"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/generator"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/publisher"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/transform"
	"github.com/mitchellh/mapstructure"
	//	logger "github.com/sirupsen/logrus"
)

type CfgMap struct {
	PipesCfgMap []PipeCfgMap `mapstructure:"pipelines"`
}

type PipeCfgMap struct {
	NameCfgMap  string      `mapstructure:"name"`
	GenCfgMap   StageCfgMap `mapstructure:"generator"`
	TransCfgMap StageCfgMap `mapstructure:"transform"`
	PubCfgMap   StageCfgMap `mapstructure:"publisher"`
}

type StageCfgMap struct {
	Type    string      `mapstructure:"type"`
	RawConf interface{} `mapstructure:"config"`
}

// ParseCfg parse the pipe array configuration
func ParseConfig(cfg interface{}) (*CfgMap, error) {
	c := &CfgMap{}
	err := mapstructure.Decode(cfg, c)
	if err != nil {
		return nil, fmt.Errorf("could not decode toplevel config: %s", err)
	}
	return c, nil
}

// ParseLinGenCfg parse the linear generator configuration
func ParseLinGenCfg(cfg interface{}) (*generator.LinearConfig, error) {
	lgc := generator.LinearConfig{}
	err := mapstructure.Decode(cfg, &lgc)
	if err != nil {
		return nil, fmt.Errorf("parselingen: could not decode linear generator config: %s", err)
	}
	return &lgc, nil
}

// ParseLinGenCfg parse the linear generator configuration
func ParseSineGenCfg(cfg interface{}) (*generator.SineConfig, error) {
	sgc := generator.SineConfig{}
	err := mapstructure.Decode(cfg, &sgc)
	if err != nil {
		return nil, fmt.Errorf("parselingen: could not decode sine generator config: %s", err)
	}
	return &sgc, nil
}

// ParseNoiseTransCfg parse the noise transformation configuration
func ParseNoiseTransCfg(cfg interface{}) (*transform.NoiseConfig, error) {
	ntc := transform.NoiseTransform{}
	err := mapstructure.Decode(cfg, &ntc)
	if err != nil {
		return nil, fmt.Errorf("parsenoisetrans: could not decode noise transform config: %s", err)
	}
	return &ntc, nil
}

// ParseSimplePubCfg parse the simple publisher configuration
func ParseSimplePubCfg(cfg interface{}) (*publisher.SimpleConfig, error) {
	spc := publisher.SimpleConfig{}
	err := mapstructure.Decode(cfg, &spc)
	if err != nil {
		return nil, fmt.Errorf("parsesimplepub: could not decode simple publisher config: %s", err)
	}
	return &ntc, nil
}

//// ParseTransCfg parse the transform configuration
//func ParseTransCfg(cfg map[string]interface{}) {}
//
//// ParsePubCfg parse the publisher configuration
//func ParsePubCfg(cfg map[string]interface{}) {}