package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/SiemensIndustrialEdgeITA/signal-generator/generator"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
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
		// Configure the data generator
		conf := &generator.Config{
			SampleRate: 1000 * time.Millisecond,
			Bufflen:    1000,
			MinVal:     0,
			MaxVal:     100,
		}
		// Create new data generator
		gen, err := generator.NewGenerator(generator.LINEAR, conf)
		if err != nil {
		}
		// Start data in parallel goroutine
		go gen.Start()

		// Stop after 5 seconds
		go func() {
			time.Sleep(5 * time.Second)
			gen.Stop()
		}()
		time.Sleep(100 * time.Second)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.signal-generator.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("user", "u", "simatic", "databus username")
	rootCmd.Flags().StringP("password", "p", "simatic", "databus password")
	rootCmd.MarkFlagRequired("user")
	rootCmd.MarkFlagRequired("password")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".signal-generator" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".signal-generator")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
