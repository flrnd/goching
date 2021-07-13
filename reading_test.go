package goching

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_hexagram_findRelatingHexagram(t *testing.T) {
	type fields struct {
		Number       int
		BinaryString string
	}
	type args struct {
		lines []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Hexagram
	}{
		{"111111", fields{1, "111111"}, args{[]int{0, 1, 2, 3, 4, 5}}, Hexagram{2, "000000"}},
		{"111111", fields{1, "111111"}, args{[]int{0, 2}}, Hexagram{6, "010111"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hex := Hexagram{
				Number:       tt.fields.Number,
				BinaryString: tt.fields.BinaryString,
			}
			if got := hex.findRelatingHexagram(tt.args.lines); !cmp.Equal(got, tt.want) {
				t.Errorf("hexagram.findRelatingHexagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toBinary(t *testing.T) {
	type args struct {
		hex []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Yin", args{[]string{"Yin", "Yin", "Yin", "Yin", "Yin", "Yin"}}, "000000"},
		{"Yang", args{[]string{"Yang", "Yang", "Yang", "Yang", "Yang", "Yang"}}, "111111"},
		{"OYin", args{[]string{"Yang", "OYin", "Yin", "Yin", "Yin", "Yin"}}, "100000"},
		{"OYang", args{[]string{"Yang", "Yin", "Yin", "Yin", "Yin", "OYang"}}, "100001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toBinary(tt.args.hex); got != tt.want {
				t.Errorf("toBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_movingLines(t *testing.T) {
	type args struct {
		hex []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"OYin", args{[]string{"Yang", "OYin", "Yin", "Yin", "Yin", "Yin"}}, []int{1}},
		{"OYin,OYang", args{[]string{"Yang", "OYin", "Yin", "OYang", "Yin", "Yin"}}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingLines(tt.args.hex); !cmp.Equal(got, tt.want) {
				t.Errorf("movingLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newCastStub() yarrow {
	return yarrow{"Yang", "OYang", "OYang", "Yin", "Yang", "Yin"}
}

func Test_CastReading(t *testing.T) {
	hex := Hexagram{5, "111010"}
	movingLines := []int{1, 2}
	resultingHex := Hexagram{3, "100010"}
	read := Reading{hex, newCastStub(), movingLines, resultingHex}

	type args struct {
		y yarrow
	}

	tests := []struct {
		name string
		args args
		want Reading
	}{
		{"1", args{yarrows}, read},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.y.CastReading(); !cmp.Equal(got, tt.want) {
				t.Errorf("CastReading() = %v, want %v", got, tt.want)
			}
		})
	}
}
