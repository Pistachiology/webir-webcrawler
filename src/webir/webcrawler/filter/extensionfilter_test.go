package filter

import (
	"reflect"
	"testing"
)

func TestExtensionFuncFactory(t *testing.T) {
	filter := ExtensionFuncFactory(nil)
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			"filter php with query",
			"http://www.google.com/file.php?query=abc.ht",
			true,
		},
		{
			"filter php without query",
			"http://www.google.com/file.php",
			true,
		},
		{
			"filter url without extension but with query",
			"http://www.google.com/?abc=q.ht",
			true,
		},
		{
			"filter url without extension, query",
			"http://www.google.com/",
			true,
		},
		{
			"filter unknown extension",
			"http://www.facebook.com/hello.dotd",
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filter(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtensionFuncFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
