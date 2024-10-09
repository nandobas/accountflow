package repository

import "accountflow/modules/entries"

type Repository interface {
	AppendEntry(entry entries.Entry) error
	GetEntriesByAccountID(accountID int64) ([]entries.Entry, error)
}
