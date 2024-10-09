package entries

import (
	"accountflow/modules/entries/repository"
	"fmt"
)

type Service interface {
	GetBalanceByAccountID(accountID int64) (float64, error)
	AppendEntry(entry Entry) error
}

func NewService(r repository.Repository) Service {
	return &entriesService{
		repository: r,
	}
}

type entriesService struct {
	repository repository.Repository
}

func (s *entriesService) GetBalanceByAccountID(accountID int64) (float64, error) {

	balance := 0.00
	entriesAccount, err := s.repository.GetEntriesByAccountID(accountID)
	if err != nil {
		return 0, fmt.Errorf("get balance by account id: unable to get balance: %w", err)
	}

	for _, givenEntry := range entriesAccount {
		switch givenEntry.EntryType {
		case EntryTypeDeposity:
			balance = balance + givenEntry.Amount

		case EntryTypeWithdrawal:
			balance = balance - givenEntry.Amount

		}

	}

	return balance, nil
}

func (s *entriesService) AppendEntry(entry Entry) error {

	return s.repository.AppendEntry(repository.Entry{
		AccountID: entry.AccountID,
		Amount:    entry.Amount,
		EntryType: entry.EntryType,
	})
}
