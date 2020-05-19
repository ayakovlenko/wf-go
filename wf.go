package wf

import (
	"fmt"
	"io"
	"wf/opml"
)

type Item struct {
	Title    string  `json:"title"`
	Note     *string `json:"note,omitempty"`
	Children []Item  `json:"items,omitempty"`
}

func Note(s string) *string {
	return &s
}

func PrintItem(w io.Writer, wfi Item) {
	printItemIndent(w, wfi, "")
}

func printItemIndent(w io.Writer, wfi Item, indent string) {
	fmt.Fprintf(w, "%s- %s\n", indent, wfi.Title)

	for _, child := range wfi.Children {
		printItemIndent(w, child, indent+"\t")
	}
}

func ToOPML(item *Item) opml.Root {
	return opml.New([]opml.Outline{
		toOutline(item),
	})
}

func toOutline(item *Item) opml.Outline {
	outlines := []opml.Outline{}
	for _, child := range item.Children {
		outlines = append(outlines, toOutline(&child))
	}

	return opml.Outline{
		Text:     item.Title,
		Note:     item.Note,
		Children: outlines,
	}
}
