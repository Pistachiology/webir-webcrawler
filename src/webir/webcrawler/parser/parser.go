package parser

import (
	"fmt"
	"regexp"
	"strings"

	"net/url"
	"path/filepath"

	"github.com/jackdanger/collectlinks"
	"golang.org/x/net/html"
)

const _parserRegex = `(http|ftp|https)://([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`

// DoCollectlinksParse all link
func DoCollectlinksParse(body string, fullURL string) []string {
	rawList := collectlinks.All(strings.NewReader(body))
	var list []string
	for _, item := range rawList {
		list = append(list, fixUrl(item, fullURL))
	}
	return list
}

func fixUrl(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}

// DoSimpleParse return list of url parse by simple regex a lot of BUG!
func DoSimpleParse(body string, fullURL string) []string {
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		return nil
	}
	var f func(*html.Node)
	var urlList []string
	baseURL, err := url.Parse(fullURL)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					var prefix, suffix string
					u := a.Val
					if !strings.HasPrefix(u, "http") {
						prefix = strings.Join([]string{baseURL.Scheme, "://", baseURL.Host, "/"}, "")
						if len(u) != 0 && u[0] == '/' {
							suffix = u[1:]
						} else if fullURL[len(fullURL)-1] == '/' {
							suffix = fmt.Sprintf("%v%v", fullURL, u)
						} else {
							suffix = fmt.Sprintf("%v/%v", filepath.Dir(baseURL.Path), u)
						}
						u = fmt.Sprintf("%v%v", prefix, suffix)
					}
					urlList = append(urlList, u)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return urlList
}

func doSimpleRegexParse(body string) []string {
	re, err := regexp.Compile(_parserRegex)
	if err != nil {
		return nil
	}
	urlList := re.FindAllString(body, -1)
	return urlList
}
