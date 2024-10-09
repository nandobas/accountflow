package storage

import "accountflow/modules/entries"

type Repository interface {
	AppendEntry(entry entries.Entry) error
	GetAllEntries() ([]entries.Entry, error)
	GetEntryByAccountID(accountID int64) (entries.Entry, error)
}
