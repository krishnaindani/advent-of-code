package main

import (
	"strings"
	"testing"
)

func Test_getColumnNumberV2(t *testing.T) {
	type args struct {
		seat string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test one",
			args: args{
				seat: "BFFFBBF",
			},
			want: 70,
		},
		{
			name: "test two",
			args: args{
				seat: "FFFBBBF",
			},
			want: 14,
		},
		{
			name: "test three",
			args: args{
				seat: "BBFFBBF",
			},
			want: 102,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getColumnNumberV2(strings.ToLower(tt.args.seat)); got != tt.want {
				t.Errorf("getColumnNumberV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSeatPositionV2(t *testing.T) {
	type args struct {
		seat string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test one",
			args: args{
				seat: "RRR",
			},
			want: 7,
		},

		{
			name: "test two",
			args: args{
				seat: "RLL",
			},
			want: 4,
		},

		{
			name: "test three",
			args: args{
				seat: "RLR",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSeatPositionV2(strings.ToLower(tt.args.seat)); got != tt.want {
				t.Errorf("getSeatPositionV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
