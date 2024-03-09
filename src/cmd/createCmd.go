package cmd

import (
	"github.com/jorgevvs2/dockeryzer/src/functions"
	"github.com/spf13/cobra"
)

var imageName string
var ignoreComments bool

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Command to generate a Dockerfile and create an Docker image (optional)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// This function will be executed when the "subcommand" is called
		functions.Create(imageName, ignoreComments)
	},
}

func init() {
	createCmd.Flags().StringVarP(&imageName, "imageName", "n", "", "Image imageName to create")
	createCmd.Flags().BoolVarP(&ignoreComments, "ignore-comments", "i", false, "No include comments to Dockerfile")

	rootCmd.AddCommand(createCmd)
}
