package validate

import (
	"testing"
)

func Test_birthYearValidation(t *testing.T) {
	type args struct {
		birthYear int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid birth year",
			args: args{
				birthYear: 1920,
			},
			want: true,
		},
		{
			name: "valid birth year",
			args: args{
				birthYear: 2002,
			},
			want: true,
		},
		{
			name: "invalid birth year",
			args: args{
				birthYear: 2004,
			},
			want: false,
		},
		{
			name: "valid birth year",
			args: args{
				birthYear: 1950,
			},
			want: true,
		},
		{
			name: "invalid birth year",
			args: args{
				birthYear: 1910,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BirthYear(tt.args.birthYear); got != tt.want {
				t.Errorf("birthYearValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heightValidation(t *testing.T) {
	type args struct {
		height string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid height in cm",
			args: args{
				height: "190cm",
			},
			want: true,
		},
		{
			name: "low invalid height in cm",
			args: args{
				height: "149",
			},
			want: false,
		},
		{
			name: "high invalid height in cm",
			args: args{
				height: "194",
			},
			want: false,
		},
		{
			name: "valid height in in",
			args: args{
				height: "75in",
			},
			want: true,
		},
		{
			name: "low invalid height in in",
			args: args{
				height: "58in",
			},
			want: false,
		},
		{
			name: "high invalid height in in",
			args: args{
				height: "77in",
			},
			want: false,
		},
		{
			name: "invalid - no symbol",
			args: args{
				height: "77",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Height(tt.args.height); got != tt.want {
				t.Errorf("heightValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_issueYearValidation(t *testing.T) {
	type args struct {
		issueYear int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy path",
			args: args{
				issueYear: 2015,
			},
			want: true,
		},
		{
			name: "greater 2020",
			args: args{
				issueYear: 2021,
			},
			want: false,
		},
		{
			name: "less then 2010",
			args: args{
				issueYear: 2009,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IssueYear(tt.args.issueYear); got != tt.want {
				t.Errorf("issueYearValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expirationYearValidation(t *testing.T) {
	type args struct {
		expirationYear int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy path",
			args: args{
				expirationYear: 2025,
			},
			want: true,
		},
		{
			name: "greater then 2030",
			args: args{
				expirationYear: 2031,
			},
			want: false,
		},
		{
			name: "less then 2020",
			args: args{
				expirationYear: 2019,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpirationYear(tt.args.expirationYear); got != tt.want {
				t.Errorf("expirationYearValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_passportIDValidation(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid passport id",
			args: args{
				id: "000000001",
			},
			want: true,
		},
		{
			name: "invalid passport id",
			args: args{
				id: "0123456789",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PassportID(tt.args.id); got != tt.want {
				t.Errorf("passportIDValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eyeColorValiadation(t *testing.T) {
	type args struct {
		color string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid eye color",
			args: args{
				color: "amb",
			},
			want: true,
		},
		{
			name: "valid eye color",
			args: args{
				color: "blu",
			},
			want: true,
		},
		{
			name: "valid eye color",
			args: args{
				color: "gry",
			},
			want: true,
		},
		{
			name: "valid eye color",
			args: args{
				color: "brn",
			},
			want: true,
		},
		{
			name: "valid eye color",
			args: args{
				color: "hzl",
			},
			want: true,
		},
		{
			name: "valid eye color",
			args: args{
				color: "oth",
			},
			want: true,
		},
		{
			name: "valid eye color",
			args: args{
				color: "grn",
			},
			want: true,
		},
		{
			name: "invalid eye color",
			args: args{
				color: "abc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EyeColor(tt.args.color); got != tt.want {
				t.Errorf("eyeColorValiadation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hairColorValidation(t *testing.T) {
	type args struct {
		color string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid hair color",
			args: args{
				color: "#123abc",
			},
			want: true,
		},
		{
			name: "invalid hair color",
			args: args{
				color: "#123abz",
			},
			want: false,
		},
		{
			name: "invalid hair color",
			args: args{
				color: "123abc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HairColor(tt.args.color); got != tt.want {
				t.Errorf("hairColorValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}
