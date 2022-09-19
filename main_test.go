package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns false for any string less than 2 characters",
			args: args{
				str: "1",
			},
			want: false,
		},
		{
			name: "returns true for a valid palindrome of odd length",
			args: args{
				str: "aba",
			},
			want: true,
		},
		{
			name: "returns true for a valid palindrome of even length",
			args: args{
				str: "aabbaa",
			},
			want: true,
		},
		{
			name: "returns true for a valid palindrome containing spaces",
			args: args{
				str: "hello world dlrow olleh",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.str); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handlePiped(t *testing.T) {
	type args struct {
		stdIn io.Reader
	}

	createFileFromString := func(s string) io.Reader {
		return strings.NewReader(s)
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "parses a piped file and returns an array of strings",
			args: args{
				stdIn: createFileFromString("aabbaa"),
			},
			want: []string{"aabbaa"},
		},
		{
			name: "parses a piped file and returns an array of new line seperated strings",
			args: args{
				stdIn: createFileFromString("aabbaa\nbbaabb\nabc"),
			},
			want: []string{"aabbaa", "bbaabb", "abc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handlePiped(tt.args.stdIn)
			if (err != nil) != tt.wantErr {
				t.Errorf("handlePiped() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handlePiped() = %v, want %v", got, tt.want)
			}
		})
	}
}
