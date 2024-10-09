package entries

const (
	EntryTypeDeposity   int = 1
	EntryTypeWithdrawal int = 2
)

type Entry struct {
	ID        int64
	AccountID int64
	Amount    float64
	EntryType int
}
