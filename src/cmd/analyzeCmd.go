package cmd

import (
	"fmt"
	"github.com/jorgevvs2/dockeryzer/src/functions"
	"github.com/spf13/cobra"
	"os"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Command to analyze a Docker image",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide an image to analyze")
			os.Exit(0)
		}

		image := args[0]
		functions.Analyze(image)
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
