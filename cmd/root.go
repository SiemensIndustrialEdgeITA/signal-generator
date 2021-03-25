package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/generator"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/pipeline"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/publisher"
	"github.com/SiemensIndustrialEdgeITA/signal-generator/transform"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "signal-generator",
	Short: `generate data stream for local testing`,
	Long:  "simple lighweight Simatic Edge App written in go which generates customizable streams of data suitable for local testing",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("config:", cfgFile)

		// Configure the data generator
		gconf := generator.LinearConfig{
			SampleRate: 1000 * time.Millisecond,
			Coeff:      0.1,
			MinVal:     0,
			MaxVal:     100,
		}

		// Configure the noise transform
		tconf := transform.NoiseConfig{
			Coeff:  0,
			MinVal: -10,
			MaxVal: 10,
		}

		// Configure the publisher mqtt client
		pconf := publisher.SimpleConfig{
			Mqtt: publisher.MqttConfig{
				Host:     "ie-databus",
				Port:     1883,
				User:     "simatic",
				Password: "simatic",
				ClientId: "signal-generator",
			},
		}

		// Create new data pipeline
		pip, _ := pipeline.NewPipeline()
		pip.AddGenerator(generator.LINEAR, gconf)
		pip.AddTransform(transform.NOISE, tconf)
		pip.AddPublisher(publisher.SIMPLE, pconf)
		pip.Build()

		// Start the pipeline
		pip.Start()

		// Runforever
		select {}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	//	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Enable me to pass cli options
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//	rootCmd.Flags().StringP("user", "u", "simatic", "databus username")
	//	rootCmd.Flags().StringP("password", "p", "simatic", "databus password")
	//	rootCmd.MarkFlagRequired("user")
	//	rootCmd.MarkFlagRequired("password")
	rootCmd.Flags().StringVar(&cfgFile, "config", "./config.yaml", "config file (default is ./config.yaml)")
}
