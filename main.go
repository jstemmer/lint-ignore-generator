package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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

	lintConfig := convert(issues)

	xml, err := lintConfig.WriteXml()
	if err != nil {
		quit("Error creating output xml: %s\n", err)
	}

	err = ioutil.WriteFile(output, xml, 0666)
	if err != nil {
		quit("Error writing output: %s\n", err)
	}
}

func convert(issues *Issues) *LintConfiguration {
	config := &LintConfiguration{}

	lintIssues := make(map[string]LintIssue)
	for _, issue := range issues.Issues {
		// Skip issues without a file
		if len(issue.Location.File) == 0 {
			continue
		}

		if _, ok := lintIssues[issue.Id]; !ok {
			lintIssue := LintIssue{
				Id:      issue.Id,
				Ignores: make([]LintIgnore, 0),
			}
			lintIssues[issue.Id] = lintIssue
		}

		lintIssue := lintIssues[issue.Id]
		lintIssue.Ignores = append(lintIssue.Ignores, LintIgnore{issue.Location.File})
		lintIssues[issue.Id] = lintIssue
	}

	config.Issues = make(LintIssues, 0, len(lintIssues))
	for _, issue := range lintIssues {
		sort.Sort(issue.Ignores)
		config.Issues = append(config.Issues, issue)
	}
	sort.Sort(config.Issues)

	return config
}

func quit(message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, message, args...)
	os.Exit(1)
}
