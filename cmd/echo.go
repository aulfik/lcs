package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	consoleReader := bufio.NewReader(os.Stdin)
	var licText string

	nameQ := "Enter Your Name: <name of author>"
	yearQ := "Enter Your Year License: <year>"
	programQ := "Program name and idea: <one line to give the program's name and a brief idea of what it does.>"
	programNameQ := "Name Program: <program>"

	if licenseType == "MIT License" {
		fmt.Println(nameQ)
		name, _ := consoleReader.ReadString('\n')
		name = strings.TrimSuffix(name, "\n")
		fmt.Println(yearQ)
		year, _ := consoleReader.ReadString('\n')
		year = strings.TrimSuffix(year, "\n")

		licText = data.GetMitText(year, name)
	}

	if licenseType == "GNU AGPLv3" {
		fmt.Println(programQ)
		program, _ := consoleReader.ReadString('\n')
		program = strings.TrimSuffix(program, "\n")
		fmt.Println(nameQ)
		name, _ := consoleReader.ReadString('\n')
		name = strings.TrimSuffix(name, "\n")
		fmt.Println(yearQ)
		year, _ := consoleReader.ReadString('\n')
		year = strings.TrimSuffix(year, "\n")

		licText = data.GetGnuAgplv3Text(program, year, name)
	}

	if licenseType == "GNU GPLv3" {
		fmt.Println(programNameQ)
		programName, _ := consoleReader.ReadString('\n')
		programName = strings.TrimSuffix(programName, "\n")
		fmt.Println(programQ)
		program, _ := consoleReader.ReadString('\n')
		program = strings.TrimSuffix(program, "\n")
		fmt.Println(nameQ)
		name, _ := consoleReader.ReadString('\n')
		name = strings.TrimSuffix(name, "\n")
		fmt.Println(yearQ)
		year, _ := consoleReader.ReadString('\n')
		year = strings.TrimSuffix(year, "\n")

		licText = data.GetGnuGplv3Text(program, year, name, programName)
	}

	if licenseType == "GNU LGPLv3" {
		licText = data.GetGnuLgplv3Text()
	}

	if licenseType == "Mozilla Public License 2.0" {
		licText = data.GetMozillaPL2Text()
	}

	if licenseType == "Apache License 2.0" {
		fmt.Println(nameQ)
		name, _ := consoleReader.ReadString('\n')
		name = strings.TrimSuffix(name, "\n")
		fmt.Println(yearQ)
		year, _ := consoleReader.ReadString('\n')
		year = strings.TrimSuffix(year, "\n")

		licText = data.GetApacheText(year, name)
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
