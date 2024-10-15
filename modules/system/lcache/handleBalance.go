package lcache

import (
	"time"

	"github.com/sirupsen/logrus"
)

func GetHandleBalanceByID(accountID int) interface{} {

	if lc, ok := SystemLocalCache.HandleBalance[accountID]; ok {

		logrus.Info("GetHandleBalanceByID from cache: ", accountID, lc.Control)

		if IsToRefresh(lc.Control) {
			return nil
		}
		return lc.Item
	}

	return nil
}

func SetHandleInfoCache(accountID int, m interface{}) {

	tn := time.Now()

	mu.Lock()

	lc := SystemLocalCache.HandleBalance[accountID]

	lc.Control.Updated = &tn
	lc.Control.ReadCount = 0
	lc.Item = m

	SystemLocalCache.HandleBalance[accountID] = lc
	mu.Unlock()
}
