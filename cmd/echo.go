package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/aulfik/lcs/data"
)

// the questions to ask
var qs = []*survey.Question{
	{
		Name: "licenseType",
		Prompt: &survey.Select{
			Message: "Choose a license:",
			Options: []string{
				"MIT License",
				"GNU AGPLv3",
				"GNU GPLv3",
				"GNU LGPLv3",
				"Mozilla Public License 2.0",
				"Apache License 2.0",
				"Boost Software License 1.0",
				"The Unlicense",
			},
			Default: "MIT License",
		},
	},
}

func generateFile(licenseType string) {
	f, err := os.Create("LICENSE")
	if err != nil {
		fmt.Printf("%v", err)
		return // It will terminate the process and not execute the rest of the codes
	}

	defer f.Close()

	var licText string

	if licenseType == "MIT License" {
		licText = data.GetMitText("2022", "Aulia")
	}

	if licenseType == "GNU AGPLv3" {
		licText = data.GetGnuAgplv3Text("Program", "2022", "Aulia Fikri")
	}

	if licenseType == "GNU GPLv3" {
		licText = data.GetGnuGplv3Text("Program Well", "2022", "Aulia Fikri", "My Program")
	}

	if licenseType == "GNU LGPLv3" {
		licText = data.GetGnuLgplv3Text()
	}

	if licenseType == "Mozilla Public License 2.0" {
		licText = data.GetMozillaPL2Text()
	}

	if licenseType == "Apache License 2.0" {
		licText = data.GetApacheText("2022", "Aulia Fikri")
	}

	if licenseType == "Boost Software License 1.0" {
		licText = data.GetBoostText()
	}

	if licenseType == "The Unlicense" {
		licText = data.GetUnlicenseText()
	}

	_, err2 := f.WriteString(licText)
	if err2 != nil {
		fmt.Printf("%v", err)
		return
	}
}
