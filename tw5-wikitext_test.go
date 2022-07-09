package main

import (
	"testing"
)

func Test_convertPairedFormatting(t *testing.T) {
	type args struct {
		txt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"bold", args{"''bold''"}, "**bold**"},
		{"italic", args{"//italic//"}, "*italic*"},
		{"underlined", args{"__underlined__"}, "<u>underlined</u>"},
		{"superscript", args{"^^superscript^^"}, "<sup>superscript</sup>"},
		{"subscript", args{",,subscript,,"}, "<sub>subscript</sub>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertPairedFormatting(tt.args.txt); got != tt.want {
				t.Errorf("convertPairedFormatting() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_tw5Tiddler_markdownText(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		tid  tw5Tiddler
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.tid.markdownText(); got != tt.want {
// 				t.Errorf("tw5Tiddler.markdownText() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_convertHeadings(t *testing.T) {
	type args struct {
		txt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Level 0", args{"No Heading"}, "No Heading\n"},
		{"Level 1", args{"! Heading"}, "# Heading\n"},
		{"Level 2", args{"!!Heading"}, "## Heading\n"},
		{"Level 3", args{"  !!!    Heading"}, "### Heading\n"},
		{"Level 4", args{"  !!!!Heading"}, "#### Heading\n"},
		{"Level 5", args{"!!!!! Heading!!"}, "##### Heading!!\n"},
		{"Level 6", args{"!!! !!! Heading"}, "### !!! Heading\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertHeadings(tt.args.txt); got != tt.want {
				t.Errorf("convertHeadings() = `%v`, want `%v`", got, tt.want)
			}
		})
	}
}

func Test_convertLinks(t *testing.T) {
	type args struct {
		txt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"no links", args{"[link]] [text|link]] [link space]] [text|link space]] ext[link]] ext[text|link]] ext[link space]] ext[text|link space]]"}, "[link]] [text|link]] [link space]] [text|link space]] ext[link]] ext[text|link]] ext[link space]] ext[text|link space]]\n"},
		{"links", args{"[[link]] [[text|link]] [[link space]] [[text|link space]] [ext[link]] [ext[text|link]] [ext[link space]] [ext[text|link space]]"}, "[link](link.md) [text](link.md) [link space](<link space.md>) [text](<link space.md>) [link](link.md) [text](link.md) [link space](<link space.md>) [text](<link space.md>)\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertLinks(tt.args.txt); got != tt.want {
				t.Errorf("convertLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertLink(t *testing.T) {
	type args struct {
		link string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"link 1", args{"[[link]]"}, "[link](link.md)"},
		{"link 2", args{"[[text|link]]"}, "[text](link.md)"},
		{"link 3", args{"[[link space]]"}, "[link space](<link space.md>)"},
		{"link 4", args{"[[text|link space]]"}, "[text](<link space.md>)"},
		{"link 5", args{"[ext[link]]"}, "[link](link.md)"},
		{"link 6", args{"[ext[text|link]]"}, "[text](link.md)"},
		{"link 7", args{"[ext[link space]]"}, "[link space](<link space.md>)"},
		{"link 8", args{"[ext[text|link space]]"}, "[text](<link space.md>)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertLink(tt.args.link); got != tt.want {
				t.Errorf("convertLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertTables(t *testing.T) {
	type args struct {
		txt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"table 1", args{"|!Hdr|!Hdr|\n|c|c|\n"}, "|Hdr|Hdr|\n|---|---|\n|c|c|\n"},
		{"table 2", args{"|!Hdr|!Hdr|  \n|c|c|  \n"}, "|Hdr|Hdr|  \n|---|---|\n|c|c|  \n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTables(tt.args.txt); got != tt.want {
				t.Errorf("convertTables() = `%v`, want `%v`", got, tt.want)
			}
		})
	}
}

func Test_convertNumberedLists(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"num list 1", args{"# one"}, "1. one\n"},
		{"num list 2", args{"## two"}, "    1. two\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertNumberedLists(tt.args.t); got != tt.want {
				t.Errorf("convertNumberedLists() = `%v`, want `%v`", got, tt.want)
			}
		})
	}
}

func Test_convertBulletedLists(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"bul list 1", args{"* one"}, "* one\n"},
		{"bul list 2", args{"** two"}, "    * two\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBulletedLists(tt.args.t); got != tt.want {
				t.Errorf("convertBulletedLists() = `%v`, want `%v`", got, tt.want)
			}
		})
	}
}

// func Test_convertWikiLinks(t *testing.T) {
// 	type args struct {
// 		txt string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		{"wikilink 1", args{" CamelCase "}, " [[CamelCase]] \n"},
// 		{"wikilink 2", args{" CamelCCase "}, " [[CamelCCase]] \n"},
// 		{"wikilink 3", args{" CamelC99 "}, " [[CamelC99]] \n"},
// 		{"wikilink 3", args{"[[Advanced Search|$:/AdvancedSearch]]"}, "[[Advanced Search|$:/AdvancedSearch]]\n"},
// 		{"~wikilink 1", args{" ~CamelCase "}, " CamelCase \n"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := convertWikiLinks(tt.args.txt); got != tt.want {
// 				t.Errorf("convertWikiLinks() = `%v`, want `%v`", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_addTitle(t *testing.T) {
	type args struct {
		txt   string
		title string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"add title", args{"Text\n[[link]]", "Title"}, "! Title\n\nText\n[[link]]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTitle(tt.args.txt, tt.args.title); got != tt.want {
				t.Errorf("addTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertTransclusion(t *testing.T) {
	type args struct {
		txt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"trans 1", args{"{{link}}"}, "{{link.md}}\n"},
		{"trans 2", args{"{{link space}}"}, "{{link space.md}}\n"},
		{"!trans 1", args{"{link space}}"}, "{link space}}\n"},
		{"!trans 2", args{"{{link space}"}, "{{link space}\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTransclusion(tt.args.txt); got != tt.want {
				t.Errorf("convertTransclusion() = `%v`, want `%v`", got, tt.want)
			}
		})
	}
}
