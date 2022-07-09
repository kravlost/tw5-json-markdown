package main

import "testing"

var tw5m tw5Tiddler

func init() {
	tw5m.Bag = "Bag"
	tw5m.Created = "20220708171600000"
	tw5m.Creator = "Creator"
	tw5m.List = "List"
	tw5m.Modified = "20230809171611111"
	tw5m.Modifier = "Modifier"
	tw5m.Revision = "Revision"
	tw5m.Tags = "one two [[three space]] four"
	tw5m.Text = "Text"
	tw5m.Title = "Title"
	tw5m.Type = "Type"
}

func Test_tw5Tiddler_yamlMetadata(t *testing.T) {
	tests := []struct {
		name string
		tid  tw5Tiddler
		want string
	}{
		{"tw5m", tw5m, `---
title: Title
date: 2022-07-08
lastmod: 2023-08-09
tags:
- one
- two
- three space
- four
author: Creator
editor: Modifier
revision: Revision
---
`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tid.yamlMetadata(); got != tt.want {
				t.Errorf("tw5Tiddler.yamlMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tw5DateToYaml(t *testing.T) {
	type args struct {
		d string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"bad date", args{"202301"}, "202301"},
		{"date 1", args{"20230105"}, "2023-01-05"},
		{"date 2", args{"20230105888888"}, "2023-01-05"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tw5DateToYaml(tt.args.d); got != tt.want {
				t.Errorf("tw5DateToYaml() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tw5TagsToYaml(t *testing.T) {
	type args struct {
		tags string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"tags 1", args{"one two [[three space]] four"}, `tags:
- one
- two
- three space
- four
`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tw5TagsToYaml(tt.args.tags); got != tt.want {
				t.Errorf("tw5TagsToYaml() = %v, want %v", got, tt.want)
			}
		})
	}
}
