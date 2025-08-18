package config

import (
	"os"
)

func SetUserCredentials() {
	os.Setenv("user")
}