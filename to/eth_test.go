package to

import (
	"testing"
)

func TestGWei(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"tGWei", args{1230000000}, 1.23},
		{"tGWei2", args{123450000000}, 123.45},
		{"tGWei3", args{12345760}, 0.01234576},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToGWei(tt.args.v); got != tt.want {
				t.Errorf("GWei() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEther(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"tEther", args{"1230000000000000000"}, 1.23},
		{"tEther2", args{"123450000000000000000"}, 123.45},
		{"tEther3", args{"12345760000000000"}, 0.01234576},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromEther(tt.want).String(); got != tt.args.v {
				t.Errorf("GWei() = %v, want %v", got, tt.want)
			}
		})
	}
}
