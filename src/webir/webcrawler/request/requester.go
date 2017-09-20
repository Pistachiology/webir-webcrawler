package request

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	log "github.com/sirupsen/logrus"
)

var client *http.Client
var reChecker []*regexp.Regexp

// Get url and return body, statusCode
func Get(url string) (string, int) {
	if reChecker == nil {
		// I need to clean code here later...
		// Maybe create function getReChecker and Iterate over string append to reChecker
		re, err := regexp.Compile(`http://www.ku.ac.th/web2012/index.php?c=adms.*`)
		if err != nil {
			log.Errorf("RegChecker Compile Error")
			return "", 404
		}
		// re2, err := regexp.Compile(`https?://www.ku.ac.th/web2012/index.php?c=adms.*`)
		// if err != nil {
		// 	return "", 404
		// }
		reChecker = []*regexp.Regexp{
			re,
			// re2,
		}
	}
	if client == nil {
		client = &http.Client{
			Timeout: time.Duration(2 * time.Second),
		}
	}
	resp, err := client.Get(url)
	if err != nil {
		return "", 404
	}
	body := resp.Body
	pb, err := ioutil.ReadAll(body)
	if err != nil {
		log.Infof("Read Error: %v", err)
		return "", 404
	}
	p := string(pb)
	// for _, re := range reChecker {
	// 	if re.MatchString(url) {
	// 		p, err = charmap.Windows874.NewDecoder().String(p)
	// 		log.Infof("match %v", url)
	// 		if err != nil {
	// 			log.Fatalf("Cannot convert %v to windows-874", url)
	// 			return "", 404
	// 		}
	// 		body = ioutil.NopCloser(body)
	// 	}
	// }

	return p, resp.StatusCode
}
