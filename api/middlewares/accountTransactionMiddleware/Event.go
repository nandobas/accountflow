package accounttransactionmiddleware

import (
	"accountflow/api/middlewares"
	accounttransactions "accountflow/modules/accountTransactions"
	"accountflow/modules/entries"
	"accountflow/modules/entries/repository"
)

type events struct {
	accountTransactionService accounttransactions.Service
	fromAccountID             int64
	toAccountID               int64
	amount                    float64
}

func NewEvent(fromAccountID, toAccountID int64, amount float64) *events {
	entryRepository := repository.NewRepositoryCache()
	entryService := entries.NewService(entryRepository)
	accountTransactionService := accounttransactions.NewAccountTransactionService(entryService)

	return &events{
		accountTransactionService: accountTransactionService,
		fromAccountID:             fromAccountID,
		toAccountID:               toAccountID,
		amount:                    amount,
	}
}

func (e *events) Deposit() *middlewares.Response_t {

	entryDeposit := entries.Entry{AccountID: e.toAccountID, Amount: e.amount, EntryType: entries.EntryTypeDeposity}
	transaction, err := e.accountTransactionService.DepositAmount(entryDeposit)
	if err != nil {
		return middlewares.RetFail("0")
	}

	return middlewares.RetOkData(formatTransactionResponse(transaction))
}

func (e *events) Withdraw() *middlewares.Response_t {

	entryWithdraw := entries.Entry{AccountID: e.fromAccountID, Amount: e.amount, EntryType: entries.EntryTypeWithdrawal}
	transaction, err := e.accountTransactionService.WithdrawAmount(entryWithdraw)
	if err != nil {
		return middlewares.RetFail("0")
	}

	return middlewares.RetOkData(formatTransactionResponse(transaction))
}

func (e *events) Transfer() *middlewares.Response_t {

	transactions, err := e.accountTransactionService.TransferAmount(e.fromAccountID, e.toAccountID, e.amount)
	if err != nil {
		return middlewares.RetFail("0")
	}

	return middlewares.RetOkData(transactions)
}

func formatTransactionResponse(t accounttransactions.Transaction) map[string]interface{} {
	return map[string]interface{}{
		t.Type.String(): map[string]interface{}{
			"id":      t.Balance.ID,
			"balance": t.Balance.Balance,
		},
	}
}
