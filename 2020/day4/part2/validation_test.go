// package main

// import (
// 	"testing"
// )

// func Test_birthYearValidation(t *testing.T) {
// 	type args struct {
// 		birthYear int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		{
// 			name: "valid birth year",
// 			args: args{
// 				birthYear: 1920,
// 			},
// 			want: true,
// 		},
// 		{
// 			name: "valid birth year",
// 			args: args{
// 				birthYear: 2002,
// 			},
// 			want: true,
// 		},
// 		{
// 			name: "invalid birth year",
// 			args: args{
// 				birthYear: 2004,
// 			},
// 			want: false,
// 		},
// 		{
// 			name: "valid birth year",
// 			args: args{
// 				birthYear: 1950,
// 			},
// 			want: true,
// 		},
// 		{
// 			name: "invalid birth year",
// 			args: args{
// 				birthYear: 1910,
// 			},
// 			want: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := birthYearValidation(tt.args.birthYear); got != tt.want {
// 				t.Errorf("birthYearValidation() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_heightValidation(t *testing.T) {
// 	type args struct {
// 		height string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		{
// 			name: "valid height in cm",
// 			args: args{
// 				height: "190cm",
// 			},
// 			want: true,
// 		},
// 		{
// 			name: "low invalid height in cm",
// 			args: args{
// 				height: "149",
// 			},
// 			want: false,
// 		},
// 		{
// 			name: "high invalid height in cm",
// 			args: args{
// 				height: "194",
// 			},
// 			want: false,
// 		},
// 		{
// 			name: "valid height in in",
// 			args: args{
// 				height: "75in",
// 			},
// 			want: true,
// 		},
// 		{
// 			name: "low invalid height in in",
// 			args: args{
// 				height: "58in",
// 			},
// 			want: false,
// 		},
// 		{
// 			name: "high invalid height in in",
// 			args: args{
// 				height: "77in",
// 			},
// 			want: false,
// 		},
// 		{
// 			name: "invalid - no symbol",
// 			args: args{
// 				height: "77",
// 			},
// 			want: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := heightValidation(tt.args.height); got != tt.want {
// 				t.Errorf("heightValidation() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
