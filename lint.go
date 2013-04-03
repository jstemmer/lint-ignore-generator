package main

import (
	"bytes"
	"encoding/xml"
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

// ReadLintXml reads the given xml and returns the Issues if it unmarshals correctly.
func ReadLintXml(data []byte) (*Issues, error) {
	issues := &Issues{}

	err := xml.Unmarshal(data, issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
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
