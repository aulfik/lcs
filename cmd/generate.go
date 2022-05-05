package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Print pass og generating file",
	Long:    `All software has versions. This is lcs`,
	Run: func(cmd *cobra.Command, args []string) {
		// the answers will be written to this struct
		answers := struct {
			LicenseType string
		}{}

		// perform the questions
		err := survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		generateFile(answers.LicenseType)

	},
}
