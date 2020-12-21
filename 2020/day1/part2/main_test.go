package main

import (
	"testing"
)

func Test_findNumbersThatSum(t *testing.T) {
	type args struct {
		nums   map[int]int
		target int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
		want2 int
	}{
		{
			name: "happy path",
			args: args{
				nums: map[int]int{
					1721: 1,
					979:  1,
					366:  1,
					299:  1,
					675:  1,
					1456: 1,
				},
				target: 2020,
			},
			want:  true,
			want1: 1721,
			want2: 299,
		},
		{
			name: "failed path",
			args: args{
				nums: map[int]int{
					1721: 1,
					979:  1,
					366:  1,
					299:  1,
					675:  1,
					1456: 1,
				},
				target: 2021,
			},
			want:  false,
			want1: 0,
			want2: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := findNumbersThatSum(tt.args.nums, tt.args.target)
			if got != tt.want {
				t.Errorf("findNumbersThatSum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findNumbersThatSum() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("findNumbersThatSum() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_findThreeNumbersThatSum(t *testing.T) {
	type args struct {
		nums   map[int]int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy path",
			args: args{
				nums: map[int]int{
					1721: 1,
					979:  1,
					366:  1,
					299:  1,
					675:  1,
					1456: 1,
				},
				target: 2020,
			},
			want: 241861950,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findThreeNumbersThatSum(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findThreeNumbersThatSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
