package entries

type Service interface {
	GetBalanceByAccountID(accountID int64) (float64, error)
	AppendEntry(entry Entry) error
}
