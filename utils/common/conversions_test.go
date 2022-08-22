package utils

import (
	"testing"
)

func TestGetMD5Hash(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hasing_checks",
			args: args{
				text: "testtext",
			},
			want: "0ea2d99c9848117666c38abce16bb43e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMD5Hash(tt.args.text); got != tt.want {
				t.Errorf("GetMD5Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertStringToInt(t *testing.T) {
	type args struct {
		numeric string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "conversion_checks_string_to_int",
			args: args{
				numeric: "1",
			},
			want: 1,
		},
		{
			name: "conversion_checks_string_to_int",
			args: args{
				numeric: "abc",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertStringToInt(tt.args.numeric); got != tt.want {
				t.Errorf("ConvertStringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
