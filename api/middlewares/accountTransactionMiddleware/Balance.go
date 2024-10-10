package accounttransactionmiddleware

import (
	"accountflow/api/middlewares"
	"accountflow/modules/entries"
	"accountflow/modules/entries/repository"
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

	data, err := b.entryService.GetBalanceByAccountID(b.accountID)
	if err != nil {
		return middlewares.RetFail("0")
	}

	return middlewares.RetOkData(data)
}
