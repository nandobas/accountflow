package accounttransactions

type TransactionType string

func (t TransactionType) String() string {
	return string(t)
}

const EntryTypeUnknown = TransactionType("unknown")
const EntryTypeOrigin = TransactionType("origin")
const EntryTypeDestination = TransactionType("destination")

type Balance struct {
	ID      int64
	Balance float64
}

type Transaction struct {
	Type    TransactionType
	Balance Balance
}
