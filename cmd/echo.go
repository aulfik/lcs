package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
var qs = []*survey.Question{
	{
		Name: "licenseType",
		Prompt: &survey.Select{
			Message: "Choose a licence:",
			Options: []string{"MIT License", "GNU AGPLv3", "GNU GPLv3", "GNU LGPLv3", "Mozilla Public License 2.0", "Apache License 2.0", "Boost Software License 1.0", "The Unlicense"},
			Default: "MIT License",
		},
	},
}

func generateFile() {
	f, err := os.Create("LICENSE")
	if err != nil {
		fmt.Printf("%v", err)
	}

	defer f.Close()

	_, err2 := f.WriteString("Hello World")

	if err2 != nil {
		fmt.Printf("%v", err)
	}
}
