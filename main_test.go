package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var (
	lintOutputFile = "lint-output.xml"
	lintConfigFile = "lint-config.xml"
)

var lintReport = Issues{
	Format: 3,
	By:     "lint 21.1",
	Issues: []Issue{
		Issue{
			Id:          "DefaultLocale",
			Severity:    "Warning",
			Message:     "Implicitly ...",
			Category:    "Correctness",
			Priority:    6,
			Summary:     "Finds ...",
			Explanation: "Calling ...",
			Url:         "http://developer.android.com/reference/java/util/Locale.html#default_locale",
			Quickfix:    "",
			Location:    Location{File: "path/to/File.class"},
		},
		Issue{
			Id:          "DefaultLocale",
			Severity:    "Warning",
			Message:     "Implicitly ...",
			Category:    "Correctness",
			Priority:    6,
			Summary:     "Finds ...",
			Explanation: "Calling ...",
			Url:         "http://developer.android.com/reference/java/util/Locale.html#default_locale",
			Quickfix:    "",
			Location:    Location{File: "path/to/AnotherFile.class"},
		},
		Issue{
			Id:          "NewApi",
			Severity:    "Error",
			Message:     "Call ...",
			Category:    "Correctness",
			Priority:    6,
			Summary:     "Finds ...",
			Explanation: "This ...",
			Url:         "",
			Quickfix:    "adt",
			Location:    Location{File: "path/to/File.class"},
		},
	},
}

var convertTestCases = []struct {
	filter   string
	expected LintConfiguration
}{
	{
		filter: "",
		expected: LintConfiguration{
			Issues: LintIssues{
				LintIssue{
					Id: "DefaultLocale",
					Ignores: LintIgnores{
						LintIgnore{"path/to/AnotherFile.class"},
						LintIgnore{"path/to/File.class"},
					},
				},
				LintIssue{
					Id: "NewApi",
					Ignores: LintIgnores{
						LintIgnore{"path/to/File.class"},
					},
				},
			},
		},
	},
	{
		filter: "path/to/AnotherFile",
		expected: LintConfiguration{
			Issues: LintIssues{
				LintIssue{
					Id: "DefaultLocale",
					Ignores: LintIgnores{
						LintIgnore{"path/to/AnotherFile.class"},
					},
				},
			},
		},
	},
}

func fixture(name string, t *testing.T) []byte {
	data, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s", name))
	if err != nil {
		t.Fatalf("Error reading fixture %s: %s", name, err)
	}
	return data
}

func TestImport(t *testing.T) {
	data := fixture(lintOutputFile, t)

	issues, err := ReadLintXml(data)
	if err != nil {
		t.Fatalf("Error importing xml: %s", err)
	}

	if issues.Format != lintReport.Format {
		t.Errorf("issues.Format == %d, want %d", issues.Format, lintReport.Format)
	}

	if issues.By != lintReport.By {
		t.Errorf("issues.By == %s, want %s", issues.By, lintReport.By)
	}

	if len(issues.Issues) != len(lintReport.Issues) {
		t.Fatalf("issues.Issues length == %d, want %d", len(issues.Issues), len(lintReport.Issues))
	}

	for i, issue := range issues.Issues {
		lintReport := lintReport.Issues[i]
		if issue != lintReport {
			t.Errorf("Issue %d ==\n%#v\nwant\n%#v", i, issue, lintReport)
		}
	}
}

func TestWrite(t *testing.T) {
	config := LintConfiguration{
		Issues: []LintIssue{
			LintIssue{
				Id: "DefaultLocale",
				Ignores: []LintIgnore{
					LintIgnore{"path/to/AnotherFile.class"},
					LintIgnore{"path/to/File.class"},
				},
			},
			LintIssue{
				Id: "NewApi",
				Ignores: []LintIgnore{
					LintIgnore{"path/to/File.class"},
				},
			},
		},
	}

	data, err := config.WriteXml()
	if err != nil {
		t.Fatalf("Error writing xml: %s", err)
	}

	expected := fixture(lintConfigFile, t)

	if string(data) != string(expected) {
		t.Fatalf("Generated xml ==\n%s\nwant\n%s\n", string(data), string(expected))
	}
}

func TestConvert(t *testing.T) {
	for i, testCase := range convertTestCases {
		in := fixture(lintOutputFile, t)

		issues, err := ReadLintXml(in)
		if err != nil {
			t.Errorf("[%d] Error parsing input: %s", i, err)
			continue
		}

		config := issues.Convert(testCase.filter)

		if len(config.Issues) != len(testCase.expected.Issues) {
			t.Errorf("[%d] Generated config issues == %d, want %d", i, len(config.Issues), len(testCase.expected.Issues))
			continue
		}

		for j, issue := range config.Issues {
			exp := testCase.expected.Issues[j]

			if issue.Id != exp.Id {
				t.Errorf("[%d][%d] Generated config issue.Id == %s want %s", i, j, issue.Id, exp.Id)
				continue
			}

			if issue.Severity != exp.Severity {
				t.Errorf("[%d][%d] Generated config issue.Severity == %s want %s", i, j, issue.Severity, exp.Severity)
				continue
			}

			if len(issue.Ignores) != len(exp.Ignores) {
				t.Errorf("[%d][%d] Generated config issue.Ignores == %d want %d", i, j, len(issue.Ignores), len(exp.Ignores))
				continue
			}

			for k, ignore := range issue.Ignores {
				if ignore.Path != exp.Ignores[k].Path {
					t.Errorf("[%d][%d][%d] Generated config ignore path == %s, want %s", i, j, k, ignore.Path, exp.Ignores[k].Path)
				}
			}
		}
	}
}
