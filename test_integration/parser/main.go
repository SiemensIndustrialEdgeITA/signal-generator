package main

import (
	"fmt"
	c2s "github.com/lumontec/config2struct"
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
        - type: noise
          config:
            coeff: 0
            min: 0
            max: 10
      sinks:
        - type: simple
          config: ''
        - type: dataservice
          config: ''
    - name: "second"
      generator: 
        type: linear
        config:
            rate_ms: 1000
            coeff: 0.1
            min: 0
            max: 100
      transforms:
        - type: noise
          config:
            coeff: 0
            min: 0
            max: 10
        - type: none
          config: ''
      sinks:
        - type: simple
          config: ''
        - type: dataservice
          config: ''

`

func main() {

	config := PipeConfig{}

	err := c2s.UnmarshalYaml([]byte(data), &config)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("configuration %#v\n", config)
	//	fmt.Printf("config.Pipelines[0].Name: %s\n", config.Pipelines[0].Name)
	//	fmt.Printf("config.Pipelines[0].Name: %d\n", config.Pipelines[0].Generator.RawConf.(LinGenConfig).Rate_ms)

}
