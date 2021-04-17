package pipelines

import (
	"github.com/SiemensIndustrialEdgeITA/signal-generator/generator"
	//	"github.com/SiemensIndustrialEdgeITA/signal-generator/publisher"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/transform"
)

type Config struct {
	PipeArrCfg []PipeConfig `mirror:"pipelines"`
}

type PipeConfig struct {
	Name      string         `mirror:"name"`
	GenCfg    DynGenConfig   `mirror:"generator,dynamic=type"`
	TransfCfg DynTransConfig `mirror:"transform,dynamic=type"`
	SinkCfg   DynSinkConfig  `mirror:"publisher,dynamic=type"`
}

type DynGenConfig struct {
	Type    string      `mirror:"type"`
	RawConf interface{} `mirror:"config"`
}

func (dg *DynGenConfig) SetDynamicType(Type string) {
	switch Type {
	case "linear":
		{
			dg.RawConf = generator.LinearConfig{}
		}
	}
}

type DynTransConfig struct {
	Type    string      `mirror:"type"`
	RawConf interface{} `mirror:"config"`
}

func (dt *DynTransConfig) SetDynamicType(Type string) {
	switch Type {
	case "none":
		{
			dt.RawConf = nil
		}
	case "noise":
		{
			dt.RawConf = transform.NoiseConfig{}
		}
	}
}

type DynSinkConfig struct {
	Type    string      `mirror:"type"`
	RawConf interface{} `mirror:"config"`
}

func (ds *DynSinkConfig) SetDynamicType(Type string) {
	switch Type {
	case "simple":
		{
			ds.RawConf = nil //publisher.SimpleConfig{}
		}
	case "dataservice":
		{
			ds.RawConf = nil
		}
	}
}
