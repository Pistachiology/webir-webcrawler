package filter

import (
	"net/http"
	"strings"
)

//ContentTypeTextHTML ...
func ContentTypeTextHTML(u string) bool {
	resp, err := http.Head(u)
	if err != nil {
		return false
	}
	ct := resp.Header.Get("Content-Type")
	return strings.Contains(ct, "text/html")
}
