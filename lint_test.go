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

var expected = Issues{
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

func TestImport(t *testing.T) {
	data := fixture(lintOutputFile, t)

	issues, err := ReadLintXml(data)
	if err != nil {
		t.Fatalf("Error importing xml: %s", err)
	}

	if issues.Format != expected.Format {
		t.Errorf("issues.Format == %d, want %d", issues.Format, expected.Format)
	}

	if issues.By != expected.By {
		t.Errorf("issues.By == %s, want %s", issues.By, expected.By)
	}

	if len(issues.Issues) != len(expected.Issues) {
		t.Fatalf("issues.Issues length == %d, want %d", len(issues.Issues), len(expected.Issues))
	}

	for i, issue := range issues.Issues {
		expectedIssue := expected.Issues[i]
		if issue != expectedIssue {
			t.Errorf("Issue %d ==\n%#v\nwant\n%#v", i, issue, expectedIssue)
		}
	}
}

func TestWrite(t *testing.T) {
	config := LintConfiguration{
		Issues: []LintIssue{
			LintIssue{
				Id:       "UnusedIds",
				Severity: "ignore",
			},
			LintIssue{
				Id:     "DefaultLocale",
				Ignore: &LintIgnore{Path: "path/to/File.class"},
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

func fixture(name string, t *testing.T) []byte {
	data, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s", name))
	if err != nil {
		t.Fatalf("Error reading fixture %s: %s", name, err)
	}
	return data
}
