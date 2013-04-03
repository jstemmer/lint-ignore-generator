package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	inputFile  string
	outputFile string
	filter     string
)

func init() {
	flag.StringVar(&inputFile, "i", "", "Input: Lint XML report")
	flag.StringVar(&outputFile, "o", "lint-config.xml", "Output: Lint configuration file")
	flag.StringVar(&filter, "f", "", "Filter by path")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if len(inputFile) == 0 {
		fmt.Fprintf(os.Stderr, "No input file specified\n\n")
		flag.Usage()
		os.Exit(1)
	}

	if len(outputFile) == 0 {
		fmt.Fprintf(os.Stderr, "No output file specified\n\n")
		flag.Usage()
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		quit("Error reading input: %s\n", err)
	}

	issues, err := ReadLintXml(data)
	if err != nil {
		quit("Error reading input xml: %s\n", err)
	}

	lintConfig := issues.Convert(filter)

	xml, err := lintConfig.WriteXml()
	if err != nil {
		quit("Error creating output xml: %s\n", err)
	}

	err = ioutil.WriteFile(outputFile, xml, 0666)
	if err != nil {
		quit("Error writing output: %s\n", err)
	}
}

func quit(message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, message, args...)
	os.Exit(1)
}
