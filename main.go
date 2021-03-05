package main

import (
	"github.com/SiemensIndustrialEdgeITA/signal-generator/cmd"
	"math/rand"
	"os"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	// TODO(lumontec): initialize singleton logger
	//	logs.InitLogs()
	//	defer logs.FlushLogs()

	command := cmd.NewRootCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
