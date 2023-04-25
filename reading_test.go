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
			if got := hex.findRelatingHexagram(tt.args.lines); !cmp.Equal(*got, tt.want) {
				t.Errorf("hexagram.findRelatingHexagram() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func Test_asBinarySeqString(t *testing.T) {
	type args struct {
		hex readingCast
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Yin", args{readingCast{"Yin", "Yin", "Yin", "Yin", "Yin", "Yin"}}, "000000"},
		{"Yang", args{readingCast{"Yang", "Yang", "Yang", "Yang", "Yang", "Yang"}}, "111111"},
		{"OYin", args{readingCast{"Yang", "OYin", "Yin", "Yin", "Yin", "Yin"}}, "100000"},
		{"OYang", args{readingCast{"Yang", "Yin", "Yin", "Yin", "Yin", "OYang"}}, "100001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.hex.asBinarySeqString(); got != tt.want {
				t.Errorf("toBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMovingLines(t *testing.T) {
	type args struct {
		hex readingCast
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"OYin", args{readingCast{"Yang", "OYin", "Yin", "Yin", "Yin", "Yin"}}, []int{1}},
		{"OYin,OYang", args{readingCast{"Yang", "OYin", "Yin", "OYang", "Yin", "Yin"}}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.hex.getMovingLines(); !cmp.Equal(got, tt.want) {
				t.Errorf("movingLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newCastStub() readingCast {
	return readingCast{"Yang", "OYang", "OYang", "Yin", "Yang", "Yin"}
}

func Test_CastReading(t *testing.T) {
	hexagram := &Hexagram{5, "111010"}
	cast := newCastStub()
	lines := cast.getMovingLines()
	resulting := &Hexagram{3, "100010"}
	read := Reading{hexagram, resulting, lines}

	type args struct {
		cast readingCast
	}

	tests := []struct {
		name string
		args args
		want Reading
	}{
		{"1", args{cast}, read},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CastReading(tt.args.cast); !cmp.Equal(*got, tt.want) {
				t.Errorf("CastReading() = %v, want %v", *got, tt.want)
			}
		})
	}
}
