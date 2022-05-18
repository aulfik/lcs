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
	Short:   "Generate a license file",
	Long:    `Generate a licence file based on your input choices`,
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
