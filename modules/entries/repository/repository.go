package repository

type Repository interface {
	AppendEntry(entry Entry) error
	GetEntriesByAccountID(accountID int64) ([]Entry, error)
}
