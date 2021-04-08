package main

type PipeConfig struct {
	Pipelines []Pipe `c2s:"pipelines"`
}

type Pipe struct {
	Name      string           `c2s:"name"`
	Generator DynGenConfig     `c2s:"generator,dynamic=type"`
	Transform []DynTransConfig `c2s:"transforms,dynamic=type"`
	Sinks     []DynSinkConfig  `c2s:"sinks,dynamic=type"`
}

type DynGenConfig struct {
	Type    string      `c2s:"type"`
	RawConf interface{} `c2s:"config"`
}

type LinGenConfig struct {
	Rate_ms int     `c2s:"rate_ms"`
	Coeff   float64 `c2s:"coeff"`
	Min     int     `c2s:"min"`
	Max     int     `c2s:"max"`
}

type DynTransConfig struct {
	Type    string      `c2s:"type"`
	RawConf interface{} `c2s:"config"`
}

type NoiseTransConfig struct {
	Coeff float64 `c2s:"coeff"`
	Min   int     `c2s:"min"`
	Max   int     `c2s:"max"`
}

type DynSinkConfig struct {
	Type    string      `c2s:"type"`
	RawConf interface{} `c2s:"config"`
}

func (dg *DynGenConfig) SetDynamicType(Type string) {
	switch Type {
	case "linear":
		{
			dg.RawConf = LinGenConfig{}
		}
	}
}

func (dt *DynTransConfig) SetDynamicType(Type string) {
	switch Type {
	case "none":
		{
			dt.RawConf = nil
		}
	case "noise":
		{
			dt.RawConf = NoiseTransConfig{}
		}
	}
}

func (ds *DynSinkConfig) SetDynamicType(Type string) {
	switch Type {
	case "simple":
		{
			ds.RawConf = nil
		}
	case "dataservice":
		{
			ds.RawConf = nil
		}
	}
}
