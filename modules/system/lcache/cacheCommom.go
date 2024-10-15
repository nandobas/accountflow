package lcache

import (
	"time"
)

const ttl = time.Hour * 2

// Test
// const cycle = time.Second * 10
// const discardTTL = time.Second * 5

// Prod
var cycle = time.Minute * 1
var discardTTL = ttl * 5

var SystemLocalCache LocalCache_t

func InitLocalCache() {

	SystemLocalCache = LocalCache_t{
		AccountEntries: map[int64]UnitCache_t{},
		HandleBalance:  map[int]UnitCache_t{},
	}

	go Cleanner()
}

// Updates cache from database after ttl be reached
func IsToRefresh(item ControlCache_t) bool {

	is := item.Updated != nil && item.Updated.Add(ttl).Before(time.Now())

	return is
}

// Cleanup cache after discardTTL be reached
func Cleanner() {

	for {
		DeleteAccountInfoItens(checkCleanUp(SystemLocalCache.AccountEntries))

		time.Sleep(cycle)
	}
}

func checkCleanUp(lci interface{}) []int64 {

	mapsToDelete := []int64{}

	for k, lc := range lci.(map[int64]UnitCache_t) {

		if lc.Control.Updated.Add(discardTTL).Before(time.Now()) {
			mapsToDelete = append(mapsToDelete, k)
		}
	}

	return mapsToDelete
}
