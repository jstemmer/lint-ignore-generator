package main

import (
	"testing"
)

var xmlFile = "fixtures/lint-output.xml"

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
	issues, err := ImportLintXML(xmlFile)
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
