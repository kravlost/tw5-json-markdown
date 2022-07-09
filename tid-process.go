package main

import (
	"fmt"
	"strings"
)

// processTiddler takes a TW5Tiddler object and converts it into
// a Markdown string.
func processTiddler(tid tw5Tiddler) (string, error) {
	fmt.Printf("Processing tiddler: %s\n", tid.Title)
	var sb strings.Builder

	sb.WriteString(tid.yamlMetadata())
	sb.WriteString("\n")
	sb.WriteString(tid.markdownText())

	return sb.String(), nil
}
