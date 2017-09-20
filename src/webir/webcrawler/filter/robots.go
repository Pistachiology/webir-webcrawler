package filter

import (
	"fmt"
	"net/url"
	"strings"
	"webir/webcrawler/request"
	wcs "webir/webcrawler/storage"

	log "github.com/sirupsen/logrus"
	rb "github.com/temoto/robotstxt"
)

// RobotsCheckFactory ...
func RobotsCheckFactory(fs *wcs.FileStorage) func(string) bool {
	return func(u string) bool {

		pu, err := url.Parse(u)
		if err != nil {
			return false
		}
		rcheck := wcs.RGet(pu.Host)
		fp := strings.Join([]string{pu.Host, "/robots.txt"}, "")

		if rcheck == wcs.NeverVisitedRobots {
			log.Infof("Checking robots for site: %v", pu.Host)
			p := fmt.Sprintf("%v://%v/robots.txt", pu.Scheme, pu.Host)
			body, sc := request.Get(p)
			if sc != 200 {
				wcs.RSet(pu.Host, wcs.RobotsNotFound)
				rcheck = wcs.RobotsNotFound
			} else {
				wcs.RSet(pu.Host, wcs.RobotsFound)
				rcheck = wcs.RobotsFound
				fs.WriteURL(p, body, true)
			}
		}

		if rcheck != wcs.RobotsFound {
			return true
		}

		b, err := fs.ReadURL(fp)
		if err != nil {
			return false
		}

		r, err := rb.FromString(b)
		if err != nil {
			return true
		}

		gr := r.FindGroup("*")
		return gr.Test(pu.EscapedPath())
	}
}
