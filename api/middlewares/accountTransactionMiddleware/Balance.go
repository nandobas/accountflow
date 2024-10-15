package accounttransactionmiddleware

import (
	"accountflow/api/middlewares"
	"accountflow/modules/entries"
	"accountflow/modules/entries/repository"
	"accountflow/modules/system/lcache"
)

type balance struct {
	entryService entries.Service
	accountID    int64
}

func NewBalance(accountID int64) *balance {
	entryRepository := repository.NewRepositoryCache()
	entryService := entries.NewService(entryRepository)

	return &balance{
		entryService: entryService,
		accountID:    accountID,
	}
}

func (b *balance) GetBalance() *middlewares.Response_t {

	if cachebalance := lcache.GetHandleBalanceByID(int(b.accountID)); cachebalance != nil {
		return middlewares.RetOkData(cachebalance)
	}

	data, err := b.entryService.GetBalanceByAccountID(b.accountID)
	if err != nil {
		return middlewares.RetFail("0")
	}

	lcache.SetHandleInfoCache(int(b.accountID), data)

	return middlewares.RetOkData(data)
}
