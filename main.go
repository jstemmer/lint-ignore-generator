package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <input.xml> <output.xml>\n", os.Args[0])
		os.Exit(1)
	}

	input, output := os.Args[1], os.Args[2]

	data, err := ioutil.ReadFile(input)
	if err != nil {
		quit("Error reading input: %s\n", err)
	}

	issues, err := ReadLintXml(data)
	if err != nil {
		quit("Error reading input xml: %s\n", err)
	}

	lintConfig := issues.Convert("")

	xml, err := lintConfig.WriteXml()
	if err != nil {
		quit("Error creating output xml: %s\n", err)
	}

	err = ioutil.WriteFile(output, xml, 0666)
	if err != nil {
		quit("Error writing output: %s\n", err)
	}
}

func quit(message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, message, args...)
	os.Exit(1)
}
