package pipelines

import (
	"github.com/SiemensIndustrialEdgeITA/signal-generator/generator"
	//	"github.com/SiemensIndustrialEdgeITA/signal-generator/publisher"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/transform"
)

type Config struct {
	PipeArrCfg []PipeConfig `c2s:"pipelines"`
}

type PipeConfig struct {
	Name      string         `c2s:"name"`
	GenCfg    DynGenConfig   `c2s:"generator,dynamic=type"`
	TransfCfg DynTransConfig `c2s:"transform,dynamic=type"`
	SinkCfg   DynSinkConfig  `c2s:"publisher,dynamic=type"`
}

type DynGenConfig struct {
	Type    string      `c2s:"type"`
	RawConf interface{} `c2s:"config"`
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
	Type    string      `c2s:"type"`
	RawConf interface{} `c2s:"config"`
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
	Type    string      `c2s:"type"`
	RawConf interface{} `c2s:"config"`
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
