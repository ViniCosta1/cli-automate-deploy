package config

import (
	"fmt"
)

func AWSConnection() bool {
	fmt.Println("Connecting to AWS Cloud...")
	return true
}

func AzureConnection() bool {
	fmt.Println("Connecting to Azure Cloud...")
	return true
}
