package goching

import "testing"

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
