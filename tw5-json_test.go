package main

import (
	_ "embed"
	"testing"
)

//go:embed tiddlers.json
var tiddlerJson string

func TestProcessJson(t *testing.T) {
	type args struct {
		inputJson string
		outputDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"proc json 1", args{tiddlerJson, "markdown"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ProcessJson(tt.args.inputJson, tt.args.outputDir); (err != nil) != tt.wantErr {
				t.Errorf("ProcessJson() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// func Test_unmarshalJson(t *testing.T) {
// 	type args struct {
// 		jsondata string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    []tw5Tiddler
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := unmarshalJson(tt.args.jsonfile)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("unmarshalJson() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("umarshalJson() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
