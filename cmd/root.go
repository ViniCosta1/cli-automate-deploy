package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dpcli",
	Short: "A CLI tool for automating deployments",
	Long:  `dpcli is a command line interface tool that helps automate deployment processes.`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Welcome to dpcli!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}
