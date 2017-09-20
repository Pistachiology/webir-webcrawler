package filter

import "regexp"

// RegexFactory ...
func RegexFactory(re *regexp.Regexp, trueIfMatch bool) func(string) bool {
	return func(u string) bool {
		return re.MatchString(u) == trueIfMatch
	}
}
