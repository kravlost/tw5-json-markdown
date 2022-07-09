package main

import (
	"fmt"
	"log"
	"strings"
)

// yamlMetadata converts the metadata in a TW5Tiddler objects into
// a YAML metadata block to go at the top of a Markdown file.
func (tid tw5Tiddler) yamlMetadata() string {
	var sb strings.Builder

	writeIfNotNull := func(tag, text string) {
		if len(text) != 0 {
			sb.WriteString(fmt.Sprintf("%s: %s\n", tag, text))
		}
	}

	// yaml header
	sb.WriteString("---\n")
	writeIfNotNull("title", tid.Title)
	writeIfNotNull("date", tw5DateToYaml(tid.Created))
	writeIfNotNull("lastmod", tw5DateToYaml(tid.Modified))
	if len(tid.Tags) != 0 {
		mdTags := tw5TagsToYaml(tid.Tags)
		sb.WriteString(mdTags)
	}
	writeIfNotNull("author", tid.Creator)
	writeIfNotNull("editor", tid.Modifier)
	writeIfNotNull("revision", tid.Revision)
	//writeIfNotNull("tw5-bag", tid.Bag)
	//writeIfNotNull("tw5-list: TODO", tid.List)
	sb.WriteString("---\n")

	return sb.String()
}

// tw5DateToYaml converts a TW5 date string to YAML format
// 20220707100549632
// yyyymmddhhmmssmmm
// to
// 2022-07-07
func tw5DateToYaml(d string) string {
	if len(d) < 8 {
		log.Printf("TW5 date `%s` too short.\n", d)
		return d
	} else {
		return d[0:4] + "-" + d[4:6] + "-" + d[6:8]
	}
}

// tw5TagsToMarkdown converts the TW5 tags format:
//   tags: one [[with space]] two
// into YAML:
//   tags:
//   - one
//   - with space
//   - two
func tw5TagsToYaml(tags string) string {
	bits := strings.Split(tags, " ")

	yamlTags := "tags:\n"

	for _, t := range bits {
		if t[0] == '[' {
			yamlTags += "- " + t[2:] + " "
		} else if ln := len(t); t[ln-1:] == "]" {
			yamlTags += t[:ln-2] + "\n"
		} else {
			yamlTags += "- " + t + "\n"
		}
	}

	return yamlTags
}
