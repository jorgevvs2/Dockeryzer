package cmd

import (
	"fmt"
	"github.com/jorgevvs2/dockeryzer/src/functions"
	"github.com/spf13/cobra"
)

var imageName string

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Command to analyze a Docker image",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide an image to analyze")
			return
		}

		image := args[0]
		// This function will be executed when the "subcommand" is called
		functions.Analyze(image)
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
