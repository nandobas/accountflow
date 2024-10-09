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

type StorageEntries struct {
	entryIdx int64
}

func (s *StorageEntries) NewEntry(accountID int64, amount float64, entryType int) Entry {
	s.entryIdx++
	entryID := s.entryIdx
	return Entry{
		ID:        entryID,
		AccountID: accountID,
		Amount:    amount,
		EntryType: entryType,
	}
}
