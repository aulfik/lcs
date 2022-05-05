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
			Message: "Choose a licence:",
			Options: []string{"MIT License", "GNU AGPLv3", "GNU GPLv3", "GNU LGPLv3", "Mozilla Public License 2.0", "Apache License 2.0", "Boost Software License 1.0", "The Unlicense"},
			Default: "MIT License",
		},
	},
}

func generateFile(licenseType string) {
	f, err := os.Create("LICENSE")
	if err != nil {
		fmt.Printf("%v", err)
	}

	defer f.Close()

	var lictext string

	if licenseType == "MIT License" {
		lictext = data.GetMitText("2021", "Aulia")
	}

	if licenseType == "GNU AGPLv3" {
		lictext = data.GetGnuAgplv3Text("Program", "2021", "Aulia Fikri")
	}

	if licenseType == "GNU GPLv3" {
		lictext = data.GetGnuGplv3Text("Program Well", "2021", "Aulia Fikri", "My Program")
	}

	if licenseType == "GNU LGPLv3" {
		lictext = data.GetGnuLgplv3Text()
	}

	if licenseType == "Mozilla Public License 2.0" {
		lictext = data.GetMozillaPL2Text()
	}

	if licenseType == "Apache License 2.0" {
		lictext = data.GetApacheText("2021", "Aulia Fikri")
	}

	if licenseType == "Boost Software License 1.0" {
		lictext = data.GetBoostText()
	}

	if licenseType == "The Unlicense" {
		lictext = data.GetUnlicenceText()
	}

	_, err2 := f.WriteString(lictext)

	if err2 != nil {
		fmt.Printf("%v", err)
	}
}
