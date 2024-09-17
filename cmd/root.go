package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cli-app",
	Short: "A CLI tool to interact with APIs",
	Long:  `This is a boilerplate to implement CLI applications using Go`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default action when no subcommands are specified
		fmt.Println("Welcome to Boilerplate CLI App! Use -h for help.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return RootCmd.Execute()
}
