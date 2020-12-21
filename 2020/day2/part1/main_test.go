package main

import (
	"testing"
)

func Test_validPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid password",
			args: args{
				password: "1-3 a: abcde",
			},
			want: true,
		},
		{
			name: "invalid password",
			args: args{
				password: "1-3 b: cdefg",
			},
			want: false,
		},
		{
			name: "valid password",
			args: args{
				password: "2-9 c: ccccccccc",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPassword(tt.args.password); got != tt.want {
				t.Errorf("validPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validPasswords(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "valid password count",
			args: args{
				data: []string{
					"1-3 a: abcde",
					"1-3 b: cdefg",
					"2-9 c: ccccccccc",
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPasswords(tt.args.data); got != tt.want {
				t.Errorf("validPasswords() = %v, want %v", got, tt.want)
			}
		})
	}
}
