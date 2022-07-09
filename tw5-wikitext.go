package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

// markdownText converts the TW5 tiddler text to Markdown and
// returns it.
func (tid tw5Tiddler) markdownText() string {
	t := tid.Text

	t = addTitle(t, tid.Title)
	//t = convertWikiLinks(t)
	t = convertNumberedLists(t)
	t = convertBulletedLists(t)
	t = convertPairedFormatting(t)
	t = convertHeadings(t)
	t = convertLinks(t)
	t = convertTransclusion(t)
	t = convertTables(t)

	return t
}

// addTitle adds the tiddler title as a top-level TW5
// heading at the top of the text block. Tiddlers
// don't contain a heading as the title is used.
func addTitle(txt, title string) string {
	txt = fmt.Sprintf("! %s\n\n%s", title, txt)

	return txt
}

// NB: TiddlyWiki seems to now deprecate automatic CamelCase linking
//     so I'm not bothering to handle it.
//     https://tiddlywiki.com/static/Tiddler%2520Title%2520Policy.html
//
// convertWikiLinks converts TW5 WikiLink
// links to TW5 [[]]-style links. It also removes the
// tilde from cancelled ~CamelCase links.
//
// CamelCase      -> [[CamelCase]]
// ~CamelCase     -> CamelCase
// but:
// [[CamelCase]]  -> [[CamelCase]]
// func convertWikiLinks(txt string) string {
// 	// Regex to match a CamelCase link.
// 	// TW5 CamelCase seems to be A-Z then at
// 	// least one a-z, then one A-Z, then
// 	// zero or more of of A-Z. a-z or 0-9
// 	re := `(?:[~])?([A-Z][a-z]+[A-Z][A-Za-z0-9]*)`

// 	m1 := regexp.MustCompile(re)

// 	converted := ""

// 	// scan over lines in the string
// 	scanner := bufio.NewScanner(strings.NewReader(txt))

// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		if m1.MatchString(line) {
// 			text := m1.ReplaceAllStringFunc(line, func(match string) string {
// 				if match[0] == '~' {
// 					return match[1:]
// 				} else {
// 					return fmt.Sprintf("[[%s]]", match)
// 				}
// 			})

// 			converted += text + "\n"
// 		} else {
// 			converted += line + "\n"
// 		}
// 	}

// 	return converted
// }

// convertNumberedLists converts TW5 numbered list
// formats to Markdown, including nested lists.
//
// # one      -> 1. one
// # two      -> 1. two
// ## three   ->     1. three
func convertNumberedLists(txt string) string {
	// Regex to match a line beginning with one or more #
	// Leading spaces will be stripped before searching
	re := `^([#]+)(\s*?)(.+)`

	m1 := regexp.MustCompile(re)

	converted := ""

	// scan over lines in the string
	scanner := bufio.NewScanner(strings.NewReader(txt))

	for scanner.Scan() {
		line := scanner.Text()
		trim := strings.TrimSpace(line)

		if s := m1.FindStringSubmatch(trim); s != nil {
			prefix := strings.Replace(s[1], "#", "    ", len(s[1])-1)
			prefix = strings.Replace(prefix, "#", "1.", 1)
			text := strings.TrimSpace(s[3])
			converted += prefix + " " + text + "\n"
		} else {
			converted += line + "\n"
		}
	}

	return converted
}

// convertBulletedLists converts TW5 bulleted list
// formats to Markdown, including nested lists.
//
// * one      -> * one
// * two      -> * two
// ** three   ->     * three
func convertBulletedLists(txt string) string {
	// Regex to match a line beginning with one or more #
	// Leading spaces will be stripped before searching
	re := `^([*]+)(\s*?)(.+)`

	m1 := regexp.MustCompile(re)

	converted := ""

	// scan over lines in the string
	scanner := bufio.NewScanner(strings.NewReader(txt))

	for scanner.Scan() {
		line := scanner.Text()
		trim := strings.TrimSpace(line)

		if s := m1.FindStringSubmatch(trim); s != nil {
			prefix := strings.Replace(s[1], "*", "    ", len(s[1])-1)
			text := strings.TrimSpace(s[3])
			converted += prefix + " " + text + "\n"
		} else {
			converted += line + "\n"
		}
	}

	return converted
}

// convertPairedFormatting converts paired formatting
// markers, such as '' to **, // to *
func convertPairedFormatting(txt string) string {
	// txt:    Tiddler text
	// tw:     TiddlyWiki search string
	// reg:    TW regex match (may be the same as tw)
	// mdPre:  The Markdown tag to replace the first match
	// mdPost: The Markdown flag to replace the second match
	findAndReplace := func(txt, tw, reg, mdPre, mdPost string) string {
		if strings.Contains(txt, tw) {
			re := `(.*?)` + reg + `(.*?)` + reg + `(.*?)`
			m1 := regexp.MustCompile(re)
			txt = m1.ReplaceAllString(txt, "$1"+mdPre+"$2"+mdPost+"$3")
		}
		return txt
	}

	txt = findAndReplace(txt, "''", `''`, "**", "**")          // bold
	txt = findAndReplace(txt, "//", `//`, "*", "*")            // italic
	txt = findAndReplace(txt, "__", `__`, "<u>", "</u>")       // underline
	txt = findAndReplace(txt, "^^", `\^\^`, "<sup>", "</sup>") // superscript
	txt = findAndReplace(txt, ",,", `,,`, "<sub>", "</sub>")   // subscript

	return txt
}

// convertHeadings converts TW5 heading markers, such as !!
// to ## and ensures correct whitespace.
func convertHeadings(txt string) string {
	// Regex to match a line beginning with one or more !
	// Leading spaces will be stripped before searching
	re := `^([!]+)(\s*?)(.+)`

	m1 := regexp.MustCompile(re)

	converted := ""

	// scan over lines in the string
	scanner := bufio.NewScanner(strings.NewReader(txt))

	for scanner.Scan() {
		line := scanner.Text()
		trim := strings.TrimSpace(line)

		if s := m1.FindStringSubmatch(trim); s != nil {
			prefix := strings.ReplaceAll(s[1], "!", "#")
			text := strings.TrimSpace(s[3])
			converted += prefix + " " + text + "\n"
		} else {
			converted += line + "\n"
		}
	}

	return converted
}

// convertLinks scans through the file for TW5 links (format [[link]]
// or [ext[link]]) to convert and passes them to convertLink()
func convertLinks(txt string) string {
	re := `(\[)(?:ext)?(\[)(.+?)(\]\])`

	m1 := regexp.MustCompile(re)

	converted := ""

	// scan over lines in the txt string
	scanner := bufio.NewScanner(strings.NewReader(txt))

	for scanner.Scan() {
		line := scanner.Text()

		line = m1.ReplaceAllStringFunc(line, convertLink)

		converted += line + "\n"
	}

	return converted
}

// convertLink converts a single link
//
// [[Link]]         	 	->  [](Link.md)
// [[Text|Link]]  		 	->  [Text](Link.md)
// [[Spaced Link]]       	->  [](<Spaced Link.md>)
// [[Text|Spaced Link]]  	->  [Text](<Spaced Link.md>)
// [ext[Link]]         	 	->  [](Link.md)
// [ext[Text|Link]]  		->  [Text](Link.md)
// [ext[Spaced Link]]    	->  [](<Spaced Link.md>)
// [ext[Text|Spaced Link]]	->  [Text](<Spaced Link.md>)
func convertLink(link string) string {
	link = strings.Trim(link, "[]")
	text := link

	if link[:4] == "ext[" {
		link = link[4:]
		text = link
	}

	if index := strings.Index(link, "|"); index != -1 {
		text = link[:index]
		link = link[index+1:]
	}

	md := "[" + text + "]("

	if strings.Contains(link, " ") {
		md += "<"
	}

	md += link

	if link[:4] != "http" {
		md += ".md"
	}

	if strings.Contains(link, " ") {
		md += ">"
	}

	md += ")"

	return md
}

// convertTransclusion scans through the file for TW5 links (format [[link]]
// or [ext[link]]) to convert and passes them to convertLink()
func convertTransclusion(txt string) string {
	re := `({{)(.*?)(}})`

	m1 := regexp.MustCompile(re)

	converted := ""

	// scan over lines in the txt string
	scanner := bufio.NewScanner(strings.NewReader(txt))

	for scanner.Scan() {
		line := scanner.Text()

		line = m1.ReplaceAllStringFunc(line, func(s string) string {
			return s[:len(s)-2] + ".md}}"
		})

		converted += line + "\n"
	}

	return converted
}

// convertTables scans through the file for table lines to convert.
// It doesn't convert table alignment.
// TODO: convert table alignment
// TODO: convert tables with no headers
func convertTables(txt string) string {

	// line beginning with |
	// line beginning with |! or | ! means a header line which must be followed by a |---|
	// type line in Markdown
	reHeader := `^\|\s*!.+\|\s*$`

	mHdr := regexp.MustCompile(reHeader)

	converted := ""

	// scan over lines in the string
	scanner := bufio.NewScanner(strings.NewReader(txt))

	for scanner.Scan() {
		line := scanner.Text()

		if mHdr.MatchString(line) {
			converted += strings.ReplaceAll(line, "!", "") + "\n"
			converted += "|"
			for c := 1; c < strings.Count(line, "|"); c++ {
				converted += "---|"
			}
			converted += "\n"
		} else {
			converted += line + "\n"
		}
	}

	return converted
}
