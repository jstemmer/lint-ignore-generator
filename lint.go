package main

import (
	"encoding/xml"
	"io/ioutil"
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

func ImportLintXML(filename string) (*Issues, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	issues := &Issues{}
	err = xml.Unmarshal(data, issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}
