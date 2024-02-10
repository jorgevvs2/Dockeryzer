package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "dockeryzer",
	Short: "An CLI application to create and analyze docker images",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// This function will be executed when the root command is called
		fmt.Println("Welcome to dockeryzer! Use --help for usage.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
