package main

// func Test_processTiddler(t *testing.T) {
// 	type args struct {
// 		tid tw5Tiddler
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    string
// 		wantErr bool
// 	}{
// 		{"proc 1", args{tw5m}, `---
// title: Title
// date: 2022-07-08
// lastmod: 2023-08-09
// tags:
// - one
// - two
// - three space
// - four
// author: Creator
// editor: Modifier
// revision: Revision
// ---

// Text
// `, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := processTiddler(tt.args.tid)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("processTiddler() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("processTiddler() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
