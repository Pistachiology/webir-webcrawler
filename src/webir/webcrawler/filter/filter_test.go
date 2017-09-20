package filter

import (
	"fmt"
	"strings"
	"testing"
	wcs "webir/webcrawler/storage"
)

func TestRobotsFilter(t *testing.T) {
	fs := wcs.NewFileStorage("tmpFolder")
	sl := strings.Split(fs.URLFilePath("http://www.github.com/abc/ddd/hhh.htm"), "/")
	fmt.Println(sl[0])

}
