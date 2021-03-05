package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// RootCmd base command for our utility
	return &cobra.Command{
		Use:   "signal-generator",
		Short: "A simple signal generation app",
		Long:  `This app allows us to generate customizable signals published on Simatic Industrial Edge topics`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("ciao")
		},
	}

}
