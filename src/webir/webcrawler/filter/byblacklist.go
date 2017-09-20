package filter

import (
	wcds "webir/webcrawler/storage"
)

func ByBlackList(u string) bool {
	return wcds.BLGet(u) == wcds.NotInBlackList
}
