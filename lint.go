package main

import (
	"bytes"
	"encoding/xml"
	"sort"
	"strings"
)

type Issues struct {
	XMLName xml.Name `xml:"issues"`
	Format  int      `xml:"format,attr"`
	By      string   `xml:"by,attr"`
	Issues  []Issue  `xml:"issue"`
}

type Issue struct {
	Id          string   `xml:"id,attr"`
	Severity    string   `xml:"severity,attr"`
	Message     string   `xml:"message,attr"`
	Category    string   `xml:"category,attr"`
	Priority    int      `xml:"priority,attr"`
	Summary     string   `xml:"summary,attr"`
	Explanation string   `xml:"explanation,attr"`
	Url         string   `xml:"url,attr"`
	Quickfix    string   `xml:"quickfix,attr"`
	Location    Location `xml:"location"`
}

type Location struct {
	File string `xml:"file,attr"`
}

type LintConfiguration struct {
	XMLName xml.Name   `xml:"lint"`
	Issues  LintIssues `xml:"issue"`
}

type LintIssue struct {
	Id       string      `xml:"id,attr"`
	Severity string      `xml:"severity,attr,omitempty"`
	Ignores  LintIgnores `xml:"ignore,omitempty"`
}

type LintIssues []LintIssue

func (l LintIssues) Len() int           { return len(l) }
func (l LintIssues) Less(i, j int) bool { return l[i].Id < l[j].Id }
func (l LintIssues) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

type LintIgnore struct {
	Path string `xml:"path,attr"`
}

type LintIgnores []LintIgnore

func (l LintIgnores) Len() int           { return len(l) }
func (l LintIgnores) Less(i, j int) bool { return l[i].Path < l[j].Path }
func (l LintIgnores) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

// ReadLintXml reads the given xml and returns the Issues if it unmarshals
// correctly.
func ReadLintXml(data []byte) (*Issues, error) {
	issues := &Issues{}

	err := xml.Unmarshal(data, issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}

// Convert creates a LintConfiguration for these issues. If a filter is given,
// only files that match the filter will be included in the LintConfiguration.
func (i *Issues) Convert(filter string) *LintConfiguration {
	config := &LintConfiguration{}

	lintIssues := make(map[string]LintIssue)
	for _, issue := range i.Issues {
		// Skip issues without a file
		if len(issue.Location.File) == 0 {
			continue
		}

		// Skip paths that do not match the filter
		if len(filter) != 0 && strings.Index(issue.Location.File, filter) == -1 {
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

// WriteXml returns the xml representation of the LintConfiguration.
func (l LintConfiguration) WriteXml() ([]byte, error) {
	var buf bytes.Buffer

	data, err := xml.MarshalIndent(l, "", "\t")
	if err != nil {
		return nil, err
	}

	// remove newline from xml.Header, because xml.MarshalIndent starts with a newline
	buf.WriteString(xml.Header[:len(xml.Header)-1])
	buf.Write(data)
	buf.WriteByte('\n')

	return buf.Bytes(), nil
}
