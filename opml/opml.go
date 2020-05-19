// Package opml provides types and methods to work with a subset of OPML
// compatible with WorkFlowy.
package opml

import "encoding/xml"

// Root represents <opml> node.
type Root struct {
	XMLName xml.Name `xml:"opml"`
	Version string   `xml:"version,attr"`
	Body    Body     `xml:"body"`
}

// New creates a new <opml> node.
func New(outlines []Outline) Root {
	return Root{
		Version: "1.0",
		Body: Body{
			Outlines: outlines,
		},
	}
}

// Body represents <body> node.
type Body struct {
	XMLName  xml.Name `xml:"body"`
	Outlines []Outline
}

// Outline represents <outline> node.
type Outline struct {
	XMLName  xml.Name  `xml:"outline"`
	Text     string    `xml:"text,attr"`
	Note     *string   `xml:"_note,attr"`
	Children []Outline `xml:"outline"`
}

// Note simply returns a reference to a string.
// A hack to make it possible to have an optional _note attribute.
func Note(s string) *string {
	return &s
}
