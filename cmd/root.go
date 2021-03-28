package cmd

import (
	"io/ioutil"
	"os"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/pipeline"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "signal-generator",
	Short: `generate data stream for local testing`,
	Long:  "simple lighweight Simatic Edge App written in go which generates customizable streams of data suitable for local testing",
	// Run the command
	Run: func(cmd *cobra.Command, args []string) {

		// Open and read file
		yamlFile, err := ioutil.ReadFile(cfgFile)

		if err != nil {
			os.Exit(1)
		}

		// Unmarshal yaml to map object
		cfgmap := make(map[interface{}]interface{})
		err = yaml.Unmarshal(yamlFile, &cfgmap)
		if err != nil {
			os.Exit(1)
		}

		// Create pipes
		pipe, err := pipeline.NewPipeArray(cfgmap)
		if err != nil {
			os.Exit(1)
		}

		// Run all the pipes in parallel
		pipe.Start()

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
