package lcache

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var mu sync.Mutex

func GetAccountEntriesInfoCache(accountID int64) interface{} {

	if lc, ok := localcache.AccountEntries[accountID]; ok {

		logrus.Info("GetAccountEntriesInfoCache from cache: ", accountID, lc.Control)

		if isToRefresh(lc.Control) {
			return nil
		}
		return lc.Item
	}

	return nil
}

func SetAccountEntriesInfoCache(accountID int64, m interface{}) {

	tn := time.Now()

	mu.Lock()

	lc := localcache.AccountEntries[accountID]

	lc.Control.Updated = &tn
	lc.Control.ReadCount = 0
	lc.Item = m

	localcache.AccountEntries[accountID] = lc
	mu.Unlock()
}

func DeleteAccountInfoItens(all []int64) {

	logrus.Info("AccountEntries total size: ", len(localcache.AccountEntries))

	mu.Lock()
	for _, mtd := range all {
		delete(localcache.AccountEntries, mtd)
	}
	mu.Unlock()
}
