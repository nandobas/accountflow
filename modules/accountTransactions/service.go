package accounttransactions

import "accountflow/modules/entries"

type Service interface {
	DeposityAmount(entry entries.Entry) (Transaction, error)
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

func (s *accounttransactionService) DeposityAmount(entry entries.Entry) (Transaction, error) {
	return Transaction{}, nil
}

func (s *accounttransactionService) WithdrawAmount(entry entries.Entry) (Transaction, error) {
	return Transaction{}, nil
}

func (s *accounttransactionService) TransferAmount(fromAccountID, toAccountID int64, amount float64) ([]Transaction, error) {
	return []Transaction{}, nil
}
