package helper

import (
	"testing"
)

func TestURLtoFilePath(t *testing.T) {
	type args struct {
		ul string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TESTCASE HERE
		{
			name: "with query",
			args: args{"http://www.google.com/abcd.php?abcd"},
			want: "www.google.com/abcd.php?abcd",
		},
		{
			name: "without query",
			args: args{"http://www.google.com/abcd.php"},
			want: "www.google.com/abcd.php",
		},
		{
			name: "Unicode test",
			args: args{"http://www.google.com/abcd.php?กกก=10"},
			want: "http://www.google.com/abcd.php?aa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLtoFilePath(tt.args.ul); got != tt.want {
				t.Errorf("URLtoFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFileExists(t *testing.T) {
	type args struct {
		fp string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFileExists(tt.args.fp); got != tt.want {
				t.Errorf("IsFileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
