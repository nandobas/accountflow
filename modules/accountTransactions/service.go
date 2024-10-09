package accounttransactions

import "accountflow/modules/entries"

type Service interface {
	WithdrawAmount(entry entries.Entry) (Transaction, error)
	TransferAmount(fromAccountID, toAccountID int64, amount float64) ([]Transaction, error)
}
