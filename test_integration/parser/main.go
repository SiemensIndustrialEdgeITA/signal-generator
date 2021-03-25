package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
	"log"
)

var data = `
pipelines:
    - name: "first"
      generator: 
        type: linear
        config:
            rate_ms: 1000
            coeff: 0.1
            min: 0
            max: 100
      transforms:
        - type: none
          config:
        - type: noise
          config:
            coeff: 0
            min: 0
            max: 10
      sinks:
        - type: simple
        - type: dataservice
`

func main() {

	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(data), &m)
	//	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Decode configuration
	d := PipeConfig{}
	mapstructure.Decode(m, &d)
	fmt.Printf("configuration %v\n", d)

	// Decode pipe array
	for _, pipemap := range d.Pipelines {
		p := Pipe{}
		mapstructure.Decode(pipemap, &p)
		fmt.Printf("pipe %v\n", p)

		// Decode generator
		g := StageConfig{}
		mapstructure.Decode(p.Generator, &g)
		fmt.Printf("generator %v\n", g)

		// Decode generator config
		gc := GenConfig{}
		mapstructure.Decode(g.RawConf, &gc)
		fmt.Printf("generator config %v\n", gc)
	}
}
