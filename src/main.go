package main

import (
	"github.com/jorgevvs2/dockeryzer/src/analyze"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "dockeryzer"}

	var name, path string

	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new image",
		Run: func(cmd *cobra.Command, args []string) {
			analyze.Analyze()
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Image Name")
	cmd.Flags().StringVarP(&path, "path", "p", "", "Image Path")

	rootCmd.AddCommand(cmd)
	rootCmd.Execute()
}
