package filter

import (
	"strings"
	wch "webir/webcrawler/helper"
)

var defaultExt = []string{".htm", ".html", ".php", ""}

// ExtensionFuncFactory return function that filter extension of urlList
func ExtensionFuncFactory(ext []string) func(string) bool {
	if ext == nil {
		ext = defaultExt
	}
	return func(u string) bool {
		valid := false
		for _, e := range ext {
			valid = valid || (strings.Compare(wch.GetURLExtension(u), e) == 0)
		}
		return valid
	}
}
