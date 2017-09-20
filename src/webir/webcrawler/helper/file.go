package helper

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// URLtoFilePath ABCD
func URLtoFilePath(ul string) string {
	if len(ul) != 0 && ul[len(ul)-1] == '/' {
		ul = fmt.Sprint(ul, "index")
	}
	u, err := url.Parse(ul)
	if err != nil {
		panic(fmt.Sprintf("Cannot parse URL to FilePath (%v)", ul))
	}
	query := ""
	rq := u.RawQuery
	if strings.Compare(rq, "") != 0 {
		query = strings.Join([]string{"?", rq}, "")
	}
	st := strings.Join([]string{u.Hostname(), u.Path, query}, "")
	return st
}

// GetURLExtension from URL ex. g.co/abc.php?q=a would return php
func GetURLExtension(ul string) string {
	if len(ul) != 0 && ul[len(ul)-1] == '/' {
		return ""
	}
	u, err := url.Parse(ul)
	if err != nil {
		panic(fmt.Sprintf("Cannot parse string to URL (%v)", ul))
	}
	st := strings.Join([]string{u.Hostname(), u.Path}, "")
	return filepath.Ext(st)
}

// IsFileExists check that filepath exists or not
func IsFileExists(fp string) bool {
	_, err := os.Stat(fp)
	return !os.IsNotExist(err)
}
