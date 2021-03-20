package main

import "testing"

func Test_getOneGroupQuestionsAnswered(t *testing.T) {
	type args struct {
		question []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test one",
			args: args{
				question: []string{"abc"},
			},
			want: 3,
		},
		{
			name: "test two",
			args: args{
				question: []string{"a", "b", "c"},
			},
			want: 0,
		},
		{
			name: "test three",
			args: args{
				question: []string{"ab", "ac"},
			},
			want: 1,
		},
		{
			name: "test four",
			args: args{
				question: []string{"a", "a", "a", "a"},
			},
			want: 1,
		},
		{
			name: "test five",
			args: args{
				question: []string{"b"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOneGroupQuestionsAnswered(tt.args.question); got != tt.want {
				t.Errorf("getOneGroupQuestionsAnswered() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getState(t *testing.T) {
	type args struct {
		states []bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test one",
			args: args{
				states: []bool{true, true, true},
			},
			want: true,
		},
		{
			name: "test one",
			args: args{
				states: []bool{true, true, false},
			},
			want: false,
		},
		{
			name: "test one",
			args: args{
				states: []bool{true, false, false},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getState(tt.args.states); got != tt.want {
				t.Errorf("getState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getQuestionsAnswered(t *testing.T) {
	type args struct {
		questions [][]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test one",
			args: args{
				questions: [][]string{
					{
						"abc",
					},
					{
						"a",
						"b",
						"c",
					},
					{
						"ab",
						"ac",
					},
					{
						"a",
						"a",
						"a",
						"a",
					},
					{
						"b",
					},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getQuestionsAnswered(tt.args.questions); got != tt.want {
				t.Errorf("getQuestionsAnswered() = %v, want %v", got, tt.want)
			}
		})
	}
}
