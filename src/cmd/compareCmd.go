package cmd

import (
	"fmt"
	"github.com/jorgevvs2/dockeryzer/src/functions"
	"github.com/spf13/cobra"
	"os"
)

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Command to compare two Docker images",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			fmt.Println("Please provide two images to compare")
			os.Exit(0)
		}

		image1, image2 := args[0], args[1]

		functions.Compare(image1, image2)
	},
}

func init() {
	rootCmd.AddCommand(compareCmd)
}
