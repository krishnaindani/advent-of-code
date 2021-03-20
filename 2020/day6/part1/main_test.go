package main

import "testing"

func Test_getOnePersonQuestionsAnswered(t *testing.T) {
	type args struct {
		question string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test one",
			args: args{
				question: "abc",
			},
			want: 3,
		},
		{
			name: "test one",
			args: args{
				question: "",
			},
			want: 0,
		},
		{
			name: "test one",
			args: args{
				question: "abac",
			},
			want: 3,
		},
		{
			name: "test one",
			args: args{
				question: "aaaa",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOnePersonQuestionsAnswered(tt.args.question); got != tt.want {
				t.Errorf("getOnePersonQuestionsAnswered() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getQuestionsAnswered(t *testing.T) {
	type args struct {
		questions []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test one",
			args: args{
				questions: []string{"abc", "abc", "abac", "aaaa", "b"},
			},
			want: 11,
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
