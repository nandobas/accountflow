package entries

const (
	EntryTypeDeposity   int = 1
	EntryTypeWithdrawal int = 2
)

type Entry struct {
	AccountID int64
	Amount    float64
	EntryType int
}
