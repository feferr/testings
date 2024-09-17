package cmd

import (
	"fmt"

	"urbanmedia/go-cli-boilerplate/business"

	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Fetches info from the API",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := business.FetchInfo()
		if err != nil {
			fmt.Printf("Error fetching info: %v\n", err)
			return
		}
		fmt.Printf("\nYour IP: %s\n", result)
	},
}

func init() {
	RootCmd.AddCommand(InfoCmd)
}
