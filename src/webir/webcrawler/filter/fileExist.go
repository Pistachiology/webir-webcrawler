package filter

import (
	wch "webir/webcrawler/helper"
)

// FileNotExists check incoming url that file exists or not
func FileNotExists(u string) bool {
	fp := wch.URLtoFilePath(u)
	return !wch.IsFileExists(fp)
}
