package goching

import (
	"testing"
)

func Test_hexagramToBinaryString(t *testing.T) {
	type args struct {
		hexagram int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"1", args{1}, "111111", false},
		{"2", args{2}, "000000", false},
		{"65", args{65}, "000000", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hexagramToBinaryString(tt.args.hexagram)
			if (err != nil) != tt.wantErr {
				t.Errorf("hexagramToBinaryString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got != tt.want) != tt.wantErr {
				t.Errorf("hexagramToBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binaryStringToHexagram(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"111111", args{"111111"}, 1, false},
		{"000000", args{"000000"}, 2, false},
		{"00010a", args{"00010a"}, -1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := binaryStringToHexagram(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("binaryStringToHexagram() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("binaryStringToHexagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBinaryString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"000000", args{"000000"}, true},
		{"111000", args{"111000"}, true},
		{"1010a0", args{"1010a0"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBinaryString(tt.args.s); got != tt.want {
				t.Errorf("isValidBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
