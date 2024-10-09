package accounttransactions

import (
	"accountflow/modules/entries"
	"fmt"
)

type Service interface {
	DepositAmount(entry entries.Entry) (Transaction, error)
	WithdrawAmount(entry entries.Entry) (Transaction, error)
	TransferAmount(fromAccountID, toAccountID int64, amount float64) ([]Transaction, error)
}

func NewAccountTransactionService(entriesService entries.Service) Service {
	return &accounttransactionService{
		entriesService: entriesService,
	}
}

type accounttransactionService struct {
	entriesService entries.Service
}

func (s *accounttransactionService) DepositAmount(givenEntry entries.Entry) (Transaction, error) {

	err := s.entriesService.AppendEntry(givenEntry)
	if err != nil {
		return Transaction{}, fmt.Errorf("deposit amount: unable to deposit: %w", err)
	}

	balanceValue, err := s.entriesService.GetBalanceByAccountID(givenEntry.AccountID)
	if err != nil {
		return Transaction{}, fmt.Errorf("unable to deposit amount: %w", err)
	}
	balance := Balance{
		givenEntry.AccountID,
		balanceValue,
	}
	return Transaction{Type: EntryTypeOrigin, Balance: balance}, nil
}

func (s *accounttransactionService) WithdrawAmount(givenEntry entries.Entry) (Transaction, error) {

	balanceValue, err := s.entriesService.GetBalanceByAccountID(givenEntry.AccountID)
	if err != nil {
		return Transaction{}, fmt.Errorf("unable to withdraw amount: %w", err)
	}
	if balanceValue < givenEntry.Amount {
		return Transaction{}, fmt.Errorf("unable to withdraw amount: unavaiable value")
	}

	err = s.entriesService.AppendEntry(givenEntry)
	if err != nil {
		return Transaction{}, fmt.Errorf("unable to withdraw: %w", err)
	}

	balanceValue, err = s.entriesService.GetBalanceByAccountID(givenEntry.AccountID)
	if err != nil {
		return Transaction{}, fmt.Errorf("unable to get balance: %w", err)
	}
	balance := Balance{
		givenEntry.AccountID,
		balanceValue,
	}
	return Transaction{Type: EntryTypeOrigin, Balance: balance}, nil
}

func (s *accounttransactionService) TransferAmount(fromAccountID, toAccountID int64, amount float64) ([]Transaction, error) {
	return []Transaction{}, nil
}
