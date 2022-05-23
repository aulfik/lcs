package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lcs",
	Short: "LCS is a very fast license generator",
	Long: `A fast and flexible license generator built with love in Go.
The complete documentation is available at http://auliafikri.engineer/posts/lcs`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
