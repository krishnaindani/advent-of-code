package main

import (
	"reflect"
	"testing"
)

func Test_convertSliceToDict(t *testing.T) {
	type args struct {
		in0 []int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		// TODO: Add test cases.
		{
			name: "convert slice to map",
			args: args{
				in0: []int{1, 2, 4, 2, 2},
			},
			want: map[int]int{
				1: 1,
				2: 3,
				4: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertSliceToDict(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertSliceToDict() = %v, want %v", got, tt.want)
			}
		})
	}
}
