package lcache

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var mu sync.Mutex

func GetAllEntries() interface{} {

	var response []UnitCache_t
	if len(SystemLocalCache.AccountEntries) > 0 {
		for _, givenEntry := range SystemLocalCache.AccountEntries {
			if IsToRefresh(givenEntry.Control) {
				continue
			}
			response = append(response, givenEntry)
		}

		return response
	}

	return nil
}

func GetAccountEntryByID(entryID int64) interface{} {

	if lc, ok := SystemLocalCache.AccountEntries[entryID]; ok {

		logrus.Info("GetAccountEntryByID from cache: ", entryID, lc.Control)

		if IsToRefresh(lc.Control) {
			return nil
		}
		return lc.Item
	}

	return nil
}

func SetAccountEntryInfoCache(entryID int64, m interface{}) {

	tn := time.Now()

	mu.Lock()

	lc := SystemLocalCache.AccountEntries[entryID]

	lc.Control.Updated = &tn
	lc.Control.ReadCount = 0
	lc.Item = m

	SystemLocalCache.AccountEntries[entryID] = lc
	mu.Unlock()
}

func DeleteAccountInfoItens(all []int64) {

	logrus.Info("AccountEntries total size: ", len(SystemLocalCache.AccountEntries))

	mu.Lock()
	for _, mtd := range all {
		delete(SystemLocalCache.AccountEntries, mtd)
	}
	mu.Unlock()
}
