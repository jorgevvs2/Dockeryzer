package cmd

import (
	"github.com/jorgevvs2/dockeryzer/src/functions"
	"github.com/spf13/cobra"
)

var imageName string

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Command to analyze a docker image",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// This function will be executed when the "subcommand" is called
		functions.Analyze(imageName)
	},
}

func init() {
	analyzeCmd.Flags().StringVarP(&imageName, "image", "i", "", "Image Name")

	rootCmd.AddCommand(analyzeCmd)
}
