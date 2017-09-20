package storage

import (
	"expvar"
)

var bls *expvar.Map

const (
	// InBlackList ...
	InBlackList = iota
	// NotInBlackList ...
	NotInBlackList = iota
)

// BLGet robots
func BLGet(url string) int64 {
	if rs == nil {
		rs = expvar.NewMap("BlackListStorage")
		return NotInBlackList
	}
	rbs := rs.Get(url)
	if rbs == nil {
		return NotInBlackList
	}
	return rbs.(*expvar.Int).Value()
}

// BLSet robots
func BLSet(url string, v int64) {
	if rs == nil {
		rs = expvar.NewMap("BlackListStorage")
	}
	rs.Add(url, v)
}
