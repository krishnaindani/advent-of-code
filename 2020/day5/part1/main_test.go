package main

import (
	"testing"
)

func Test_getColumnNumber(t *testing.T) {
	type args struct {
		seat string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy path",
			args: args{
				seat: "fbfbbff",
			},
			want: 44,
		},
		{
			name: "second seat",
			args: args{
				seat: "bfffbbf",
			},
			want: 70,
		},
		{
			name: "third seat",
			args: args{
				seat: "fffbbbf",
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getColumnNumber(tt.args.seat); got != tt.want {
				t.Errorf("getColumnNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSeatPosition(t *testing.T) {
	type args struct {
		seat string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy path",
			args: args{
				seat: "rlr",
			},
			want: 5,
		},
		{
			name: "second seat",
			args: args{
				seat: "rrr",
			},
			want: 7,
		},
		{
			name: "third seat",
			args: args{
				seat: "rll",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSeatPosition(tt.args.seat); got != tt.want {
				t.Errorf("getSeatPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
