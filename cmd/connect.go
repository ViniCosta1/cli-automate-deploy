package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vinicosta1/cli-automate-deploy/config"
)

var configUserCmd = &cobra.Command{
	Short: "Configurating User Credentials",
	Long:  `This command allows you to configure your user credentials for the cloud provider.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Configuring user credentials...")
		config.SetUserCredentials()
	},
}