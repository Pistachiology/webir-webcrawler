package main

import (
	"os"
	"regexp"
	"time"
	wcds "webir/webcrawler/datastructure"
	wcf "webir/webcrawler/filter"
	wch "webir/webcrawler/helper"
	wcp "webir/webcrawler/parser"
	wcr "webir/webcrawler/request"
	wcs "webir/webcrawler/storage"

	log "github.com/sirupsen/logrus"
)

func filterListChain(list []string, filters ...func(list []string) []string) []string {
	for _, filter := range filters {
		// log.Infof("Probe: %v %v", len(list), filter)
		list = filter(list)
	}
	return list
}

var htmlCount = 0

func fileCountWatcher(fs wcs.Storage, d time.Duration) {
	ticker := time.Tick(time.Millisecond * d)
	for _ = range ticker {
		log.Infof("Crawled [%v] website [%v] are htm, html", fs.GetTotalFileCount(), htmlCount)
		if htmlCount >= 10000 {
			panic("Finish :D")
		}
	}
}

func main() {
	os.RemoveAll("html")
	fs := wcs.NewFileStorage("html")
	rm := wcr.NewManager(15, wcp.DoCollectlinksParse, fs)
	go fileCountWatcher(fs, 2000)
	// initailize
	rm.PushQueue(
		// "https://mike.cpe.ku.ac.th/01204453/home.php",
		// "http://www.ku.ac.th/king10.html",
		"http://www.ku.ac.th/web2012/index.php?c=adms&m=mainpage1",
	// "http://mirror1.ku.ac.th/",
	)
	reFilterByDomainKU, err := regexp.Compile(`^https?://[a-zA-Z0-9\.]*\.ku.ac.th(/.*)?$`)
	if err != nil {
		panic("Filter domain ku is not working !")
	}
	// reNotKUWeb2012, err2 := regexp.Compile(`https?://www.ku.ac.th/web2012/index.php.*`)
	// if err2 != nil {
	// 	panic("Filter domain ku web2012 is not working !")
	// }
	outputChan := rm.Exec(2000)
	for output := range outputChan {
		fs.WriteURL(output.URL, output.Body, true)
		if wcf.ExtensionFuncFactory([]string{".htm", ".html"})(output.URL) {
			htmlCount++
		}
		// We can make filter and pushtoQueue as thread to improve perf
		go func(output *wcds.Node) {
			edges := filterListChain(output.Edges,
				wch.FilterList(wcf.RegexFactory(reFilterByDomainKU, true)),
				wch.FilterList(wcf.ExtensionFuncFactory([]string{".php", ".htm", ".html", "", ".xml", ".asp"})),
				// wch.FilterList(wcf.RegexFactory(reNotKUWeb2012, false)),
				wcf.RemoveDuplicate,
				wch.FilterList(wcf.FileNotExists),
				// wch.FilterList(wcf.ContentTypeTextHTML),
				wch.FilterList(wcf.RobotsCheckFactory(fs)),
			)
			rm.PushQueue(edges...)
		}(output)
	}
}

/*
	Architecture
								 FileStorage
								/			^
							/				| (FilterStorage Check file Exists)
						/					v
	RequestManager --> ListFilter
	(with Parser)

	Filter: extensionFilter -> filterVisitedWebSite(by file) -> robotsCheck -> removeDuplicate (always last O(n^2))

	Inside
	RequestManager

	precallFilter -> request -> postcallFilter (TODO)
*/
