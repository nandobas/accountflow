package accounttransactions

type TransactionType string

func (t TransactionType) String() string {
	return string(t)
}

const EntryTypeUnknown = TransactionType("unknown")
const EntryTypeOrigin = TransactionType("origin")
const EntryTypeDestination = TransactionType("destination")

type Balance struct {
	ID      int64   `json:"id"`
	Balance float64 `json:"balance"`
}

type Transaction struct {
	Type    TransactionType
	Balance Balance
}
