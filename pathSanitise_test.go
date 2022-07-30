package main

import "testing"

func Test_sanitiseLeafName(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"filename", args{"f?:i/le>  n%a|m\\e.ext"}, "f__i_le_  n_a_m_e.ext"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sanitiseLeafName(tt.args.filename); got != tt.want {
				t.Errorf("sanitiseLeafName() = %v, want %v", got, tt.want)
			}
		})
	}
}
