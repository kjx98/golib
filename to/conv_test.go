package to

import (
	"testing"
)

func TestString(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"tString1", args{1.03}, "1.03"},
		{"tString2", args{1.03e55}, "1.03e+55"},
		{"tString3", args{float32(1.1234567890123456)}, "1.1234568"},
		{"tString4", args{1.12345678901234512345}, "1.123456789012345"},
		{"tString5", args{123}, "123"},
		{"tString6", args{"Hello World"}, "Hello World"},
		{"tString5", args{123456789012345}, "123456789012345"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.v); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"tIntl1", args{12345.76}, 12345},
		{"tIntl2", args{"12345"}, 12345},
		{"tIntl3", args{-12345.76}, -12345},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64(tt.args.v); got != tt.want {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"tInt1", args{12345.76}, 12345},
		{"tInt2", args{"12345"}, 12345},
		{"tInt3", args{-12345.76}, -12345},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.v); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
		{"tIntu1", args{12345.76}, 12345},
		{"tIntu2", args{"12345"}, 12345},
		{"tIntu3", args{123456789012345}, 123456789012345},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64(tt.args.v); got != tt.want {
				t.Errorf("Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
