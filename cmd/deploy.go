package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vinicosta1/cli-automate-deploy/config"
)

var deployCmd = &cobra.Command{
	Use:   "deploy <cloud_provider>",
	Short: "Deploy your application",
	Long:  `Deploy your application to the specified environment.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch cloud := args[0]; cloud {
		case "aws":
			if config.AWSConnection() {
				fmt.Println("Deploying to AWS...")
			}
		case "azure":
			if config.AzureConnection() {
				fmt.Println("Deploying to Azure...")
			}
		default:
			fmt.Println("Unknown cloud provider.")
		}
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
