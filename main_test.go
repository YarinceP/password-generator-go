package main

import "testing"

func Test_containsUpperCase(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "contains upperCase",
			args: args{
				s: []byte("Hello World"),
			},
			want: true,
		},
		{
			name: "contains lowerCase",
			args: args{
				s: []byte("hello world"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsUpperCase(tt.args.s); got != tt.want {
				t.Errorf("containsUpperCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsLowerCase(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "contains lowerCase",
			args: args{
				s: []byte("Hello World"),
			},
			want: true,
		},
		{
			name: "contains only upperCase",
			args: args{
				s: []byte("HELLO WORLD"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsLowerCase(tt.args.s); got != tt.want {
				t.Errorf("containsLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsNumbers(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "contains numbers",
			args: args{
				s: []byte("Hello World 9"),
			},
			want: true,
		},
		{
			name: "not contains numbers",
			args: args{
				s: []byte("Hello World"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNumbers(tt.args.s); got != tt.want {
				t.Errorf("containsNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_passwordGenerator(t *testing.T) {
	type args struct {
		passwordLength int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Generate password",
			args: args{
				passwordLength: 8,
			},
			want: "123456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := passwordGenerator(tt.args.passwordLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("passwordGenerator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("passwordGenerator() got = %v, want %v", got, tt.want)
			}
		})
	}
}
