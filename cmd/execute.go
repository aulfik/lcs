package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lcs",
	Short: "Lcs is a very fast licence generator",
	Long: `A Fast and Flexible Licence Generator built with
				  love in Go.
				  Complete documentation is available at http://auliafikri.engineer/lcs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
