package main

import (
	"reflect"
	"testing"
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
		want   hexagram
	}{
		{"111111", fields{1, "111111"}, args{[]int{0, 1, 2, 3, 4, 5}}, hexagram{2, "000000"}},
		{"111111", fields{1, "111111"}, args{[]int{0, 2}}, hexagram{6, "010111"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hex := hexagram{
				Number:       tt.fields.Number,
				BinaryString: tt.fields.BinaryString,
			}
			if got := hex.findRelatingHexagram(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hexagram.findRelatingHexagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func shuffleStub() []string {
	yarrow := []string{"Yin", "OYang", "OYang", "OYang", "Yang", "Yang", "Yang", "Yang", "Yang",
		"Yin", "Yin", "Yin", "Yin", "Yin", "Yin", "OYin",
	}
	return yarrow
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
			if got := movingLines(tt.args.hex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("movingLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func generateCastStub() []string {
	result := []string{"OYang", "Yin", "Yin", "OYang", "OYang", "Yang"}
	return result
}

func Test_generateReading(t *testing.T) {
	yarrow := []string{"OYin", "OYang", "OYang", "OYang", "Yang", "Yang", "Yang", "Yang", "Yang",
		"Yin", "Yin", "Yin", "Yin", "Yin", "Yin", "Yin",
	}
	hex := hexagram{25, "100111"}
	movingLines := []int{0, 3, 4}
	resultingHex := hexagram{23, "000001"}
	read := reading{hex, generateCastStub(), movingLines, resultingHex}

	type args struct {
		yarrow []string
	}
	tests := []struct {
		name string
		args args
		want reading
	}{
		{"111111", args{yarrow}, read},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateReading(tt.args.yarrow); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("castReading() = %v, want %v", got, tt.want)
			}
		})
	}
}
