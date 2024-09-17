package main

import (
	"fmt"

	"urbanmedia/go-cli-boilerplate/cmd"
	"urbanmedia/go-cli-boilerplate/config"
)

func main() {
	// Initialize configuration
	config.InitConfig()

	// Use configuration
	fmt.Println("App Name:", config.AppConfig.AppName)

	// Execute CLI commands
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
