package cmd

import (
	"github.com/jorgevvs2/dockeryzer/src/functions"
	"github.com/spf13/cobra"
)

var name string
var ignoreComments bool

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Command to generate a Dockerfile and create an Docker image (optional)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// This function will be executed when the "subcommand" is called
		functions.Create(name, ignoreComments)
	},
}

func init() {
	createCmd.Flags().StringVarP(&name, "name", "n", "", "Image name to create")
	createCmd.Flags().BoolVarP(&ignoreComments, "ignore-comments", "i", false, "No include comments to Dockerfile")

	rootCmd.AddCommand(createCmd)
}
