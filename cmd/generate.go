package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Print pass og generating file",
	Long:  `All software has versions. This is lcs`,
	Run: func(cmd *cobra.Command, args []string) {
		generateFile()
	},
}
