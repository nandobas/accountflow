package repository

type Entry struct {
	ID        int64
	AccountID int64
	Amount    float64
	EntryType int
}
