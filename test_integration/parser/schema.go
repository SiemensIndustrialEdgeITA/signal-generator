package main

type PipeConfig struct {
	Pipelines []Pipe `mirror:"pipelines"`
}

type Pipe struct {
	Name      string           `mirror:"name"`
	Generator DynGenConfig     `mirror:"generator,dynamic=type"`
	Transform []DynTransConfig `mirror:"transforms,dynamic=type"`
	Sinks     []DynSinkConfig  `mirror:"sinks,dynamic=type"`
}

type DynGenConfig struct {
	Type    string      `mirror:"type"`
	RawConf interface{} `mirror:"config"`
}

type LinGenConfig struct {
	Rate_ms int     `mirror:"rate_ms"`
	Coeff   float64 `mirror:"coeff"`
	Min     int     `mirror:"min"`
	Max     int     `mirror:"max"`
}

type DynTransConfig struct {
	Type    string      `mirror:"type"`
	RawConf interface{} `mirror:"config"`
}

type NoiseTransConfig struct {
	Coeff float64 `mirror:"coeff"`
	Min   int     `mirror:"min"`
	Max   int     `mirror:"max"`
}

type DynSinkConfig struct {
	Type    string      `mirror:"type"`
	RawConf interface{} `mirror:"config"`
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
