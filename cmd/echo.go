package cmd

import (
	"fmt"
	"os"
)

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
