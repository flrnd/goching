package dictionary

import (
	"testing"
)

func Test_GetBinaryString(t *testing.T) {
	type args struct {
		hexagram int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{1}, "111111"},
		{"2", args{2}, "000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBinaryString(tt.args.hexagram)
			if err != nil {
				t.Fatalf("GetBinaryString() error = %v", err)
			}
			if *got != tt.want {
				t.Errorf("GetBinaryString() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func Test_GetBinaryString_ErrorOnInvalidHexagramNumber(t *testing.T) {
	type args struct {
		hexagram int
	}

	tests := []struct {
		name string
		args args
		want error
	}{
		{"65", args{65}, ErrInvalidHexagramNumber},
		{"-1", args{-1}, ErrInvalidHexagramNumber},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetBinaryString(tt.args.hexagram)
			if err == nil {
				t.Errorf("GetBinaryString() error: %v, want: %v", err, tt.want)
			}
		})
	}

}
func Test_GetHexagram(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"111111", args{"111111"}, 1},
		{"000000", args{"000000"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHexagram(tt.args.s)
			if err != nil {
				t.Fatalf("GetHexagram() error = %v", err)
			}
			if *got != tt.want {
				t.Errorf("GetHexagram() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func Test_GetHexagram_ErrorOnInvalidBinaryString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{"#1as111", args{"#1as111"}, ErrInvalidBinaryString},
		{"009000", args{"009000"}, ErrInvalidBinaryString},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetHexagram(tt.args.s)
			if err == nil {
				t.Errorf("GetHexagram() error = %v, want = %v", err, tt.want)
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
