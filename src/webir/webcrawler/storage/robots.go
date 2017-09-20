package storage

import (
	"expvar"
)

var rs *expvar.Map

const (
	// NeverVisitedRobots ....
	NeverVisitedRobots = 0
	// RobotsNotFound ...
	RobotsNotFound = 1
	// RobotsFound ...
	RobotsFound = 2
)

// RGet robots
func RGet(url string) int64 {
	if rs == nil {
		rs = expvar.NewMap("RobotsStorage")
		return NeverVisitedRobots
	}
	rbs := rs.Get(url)
	if rbs == nil {
		return NeverVisitedRobots
	}
	return rbs.(*expvar.Int).Value()
}

// RSet robots
func RSet(url string, v int64) {
	if rs == nil {
		rs = expvar.NewMap("RobotsStorage")
	}
	rs.Add(url, v)
}
