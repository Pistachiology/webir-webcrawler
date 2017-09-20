package filter

import (
	"reflect"
	"regexp"
	"testing"
)

func TestRegexKUDomain(t *testing.T) {
	reFilterByDomainKU, err := regexp.Compile(`^https?://[a-zA-Z0-9\.]*\.ku.ac.th(/.*)?$`)
	if err != nil {
		t.Errorf("Cannot create regexp")
	}
	regex := RegexFactory(reFilterByDomainKU, true)
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			"simple ku domain",
			"http://www.ku.ac.th/",
			true,
		},
		{
			"ku domain with page and query",
			"https://www.ku.ac.th/index.php?lang=abbc",
			true,
		},
		{
			"ku with long subdomain",
			"https://www.qqq.ddd.ddd.ku.ac.th",
			true,
		},
		{
			"not ku domain",
			"http://google.com/",
			false,
		},
		{
			"not ku domain but ku domain in url",
			"http://www.google.com/?q=http://www.ku.ac.th",
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := regex(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestRegexKUDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
