package utils

import "testing"

func TestStrConvertUint(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			"success convert",
			args{"12"},
			12,
		},
		{
			"failure convert",
			args{"asd"},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrConvertUint(tt.args.str); got != tt.want {
				t.Errorf("StrConvertUint() = %v, want %v", got, tt.want)
			}
		})
	}
}
