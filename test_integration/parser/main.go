package main

import (
	"fmt"
	mirror "github.com/lumontec/mirror"
	//	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {

	cfgFile := "./config.yml"
	config := PipeConfig{}

	// Open and read file
	absPath, _ := filepath.Abs(cfgFile)
	yamlFile, err := ioutil.ReadFile(absPath)

	if err != nil {

		fmt.Println("could not read file:", cfgFile, " err:", err)
		os.Exit(1)
	}

	//	rawmap := make(map[string]interface{})
	//
	//	err = yaml.Unmarshal(yamlFile, rawmap)
	//	if err != nil {
	//		return
	//	}
	//
	//	fmt.Printf("unmarshallled: %#v", rawmap)

	err = mirror.UnmarshalYaml(yamlFile, &config)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("configuration %#v\n", config)
	//	fmt.Printf("config.Pipelines[0].Name: %s\n", config.Pipelines[0].Name)
	//	fmt.Printf("config.Pipelines[0].Name: %d\n", config.Pipelines[0].Generator.RawConf.(LinGenConfig).Rate_ms)

}
