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
#      transforms:
#        - type: none
#          config:
#        - type: noise
#          config:
#            coeff: 0
#            min: 0
#            max: 10
#      sinks:
#        - type: simple
#        - type: dataservice
`

func main() {
	//	m := make(map[interface{}]interface{})
	d := PipeConfig{}

	//	err := yaml.Unmarshal([]byte(data), &m)
	err := yaml.Unmarshal([]byte(data), &d)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	//	fmt.Printf("%v\n", m)
	fmt.Printf("name: %s\n", d.Pipelines[0].Name)
	fmt.Printf("type: %s\n", d.Pipelines[0].Generator.Type)
	fmt.Printf("%v\n", d.Pipelines[0].Generator.RawConf)

	genconfig := GenConfig{}
	mapstructure.Decode(d.Pipelines[0].Generator.RawConf, &genconfig)

	fmt.Printf("genconfig: %v\n", genconfig)

}

//func ParseGenerator(genmap map[string]interface{}) GenConfig {
//	result := GenConfig{}
//	mapstructure.Decode(genmap, &result)
//	return result
//}
