package request

import (
	"errors"
	"fmt"
	"sync"
	"time"
	ds "webir/webcrawler/datastructure"
	wcs "webir/webcrawler/storage"

	log "github.com/sirupsen/logrus"
)

// Manager struct
type Manager struct {
	queue    []string
	poolSize int
	parser   func(string, string) []string
	s        wcs.Storage
	sync.Mutex
}

// NewManager return Manager with proper allocation
func NewManager(poolSize int, parser func(string, string) []string, s wcs.Storage) *Manager {
	var rm Manager
	rm.poolSize = poolSize
	rm.queue = make([]string, 0)
	rm.parser = parser
	rm.s = s
	return &rm
}

// Exec run until queue is empty for n millisecond
func (rm *Manager) Exec(millisecond time.Duration) chan *ds.Node {
	workingChance := 30
	output := make(chan *ds.Node, rm.poolSize*10)
	jobs := make(chan string, rm.poolSize)

	// pool manager
	go func() {
		for failCount := 0; failCount != workingChance; {
			if queue, err := rm.getQueue(); err == nil {
				if failCount != 0 {
					log.Infoln("Found incoming work :D")
				}
				failCount = 0
				jobs <- queue
			} else {
				failCount++
				log.Infoln(fmt.Sprintf("No work found.. try waiting...chance (%v/%v)", failCount, workingChance))
				time.Sleep(time.Millisecond * millisecond)
			}
		}
		close(jobs)
		close(output)
	}()

	// pool worker
	for i := 0; i < rm.poolSize; i++ {
		go func() {
			for job := range jobs {
				if rm.s.URLExists(job) {
					log.Debugf("FILE EXISTS %v\n", job)
					continue
				}
				body, statusCode := Get(job)
				if statusCode == 404 {
					continue
				}
				edges := rm.parser(body, job)
				output <- &ds.Node{URL: job, Body: body, Edges: edges, StatusCode: statusCode}
			}
		}()
	}

	return output
}

func (rm *Manager) getQueue() (string, error) {
	rm.Lock()
	defer rm.Unlock()
	if len(rm.queue) == 0 {
		return "", errors.New("queue is empty")
	}
	ret, queue := rm.queue[0], rm.queue[1:]
	rm.queue = queue
	return ret, nil
}

// PushQueue arg urlList
func (rm *Manager) PushQueue(urlList ...string) {
	rm.Lock()
	defer rm.Unlock()
	rm.queue = append(rm.queue, urlList...)
}
