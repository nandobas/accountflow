package accounttransactions_test

import (
	accounttransactions "accountflow/modules/accountTransactions"
	"accountflow/modules/entries"
	"accountflow/modules/entries/repository"
	"accountflow/modules/system/lcache"
	"testing"

	"github.com/stretchr/testify/suite"
)

type testServiceSuite struct {
	suite.Suite
}

func TestService(t *testing.T) {
	suite.Run(t, new(testServiceSuite))
}

func (t *testServiceSuite) SetupTest() {

	lcache.InitLocalCache()
}

func (t *testServiceSuite) TestAccountTransactionService_WhenDeposityAmount_ExpectedTransactionAmount() {
	// Arrange
	accountID := int64(100)
	expectedBalance := float64(15.00)
	expectedType := accounttransactions.EntryTypeOrigin
	entryRepository := repository.NewRepositoryCache()
	entryService := entries.NewService(entryRepository)
	accountTransactionService := accounttransactions.NewAccountTransactionService(entryService)

	entry := entries.Entry{AccountID: accountID, Amount: 15.00, EntryType: entries.EntryTypeDeposity}

	// Act
	transaction, err := accountTransactionService.DepositAmount(entry)

	// Assert
	t.NoError(err)
	t.Equal(accountID, transaction.Balance.ID)
	t.Equal(expectedBalance, transaction.Balance.Balance)
	t.Equal(expectedType, transaction.Type)
}

func (t *testServiceSuite) TestAccountTransactionService_WhenWithdrawAmount_ExpectedFailUnavaiableValue() {
	// Arrange
	accountID := int64(100)
	expectedError := "unable to withdraw amount: unavaiable value"
	entryRepository := repository.NewRepositoryCache()
	entryService := entries.NewService(entryRepository)
	accountTransactionService := accounttransactions.NewAccountTransactionService(entryService)

	// deposit
	entryDeposit := entries.Entry{AccountID: accountID, Amount: 10.00, EntryType: entries.EntryTypeDeposity}
	_, err := accountTransactionService.DepositAmount(entryDeposit)
	t.NoError(err)

	// withdrawal
	entryWithdrawal := entries.Entry{AccountID: accountID, Amount: 15.00, EntryType: entries.EntryTypeWithdrawal}

	// Act
	_, err = accountTransactionService.WithdrawAmount(entryWithdrawal)

	// Assert
	t.Error(err)
	t.EqualError(err, expectedError)
}

func (t *testServiceSuite) TestAccountTransactionService_WhenWithdrawAmount_ExpectedTransactionAmount() {
	// Arrange
	accountID := int64(100)
	expectedBalance := float64(1.00)
	expectedType := accounttransactions.EntryTypeOrigin
	entryRepository := repository.NewRepositoryCache()
	entryService := entries.NewService(entryRepository)
	accountTransactionService := accounttransactions.NewAccountTransactionService(entryService)

	// deposit
	entryDeposit := entries.Entry{AccountID: accountID, Amount: 10.00, EntryType: entries.EntryTypeDeposity}
	_, err := accountTransactionService.DepositAmount(entryDeposit)
	t.NoError(err)

	// withdrawal
	entryWithdrawal := entries.Entry{AccountID: accountID, Amount: 9.00, EntryType: entries.EntryTypeWithdrawal}

	// Act
	transaction, err := accountTransactionService.WithdrawAmount(entryWithdrawal)

	// Assert
	t.NoError(err)
	t.Equal(accountID, transaction.Balance.ID)
	t.Equal(expectedBalance, transaction.Balance.Balance)
	t.Equal(expectedType, transaction.Type)
}

func (t *testServiceSuite) TestAccountTransactionService_WhenTransferAmount_ExpectedTransactionsAmount() {
	// Arrange
	var err error
	accountID_A := int64(100)
	accountID_B := int64(300)
	expectedBalanceA := float64(0.00)
	expectedBalanceB := float64(15.00)
	transferAmount := float64(15.00)
	expectedTypeA := accounttransactions.EntryTypeOrigin
	expectedTypeB := accounttransactions.EntryTypeDestination
	entryRepository := repository.NewRepositoryCache()
	entryService := entries.NewService(entryRepository)
	accountTransactionService := accounttransactions.NewAccountTransactionService(entryService)

	// deposit A
	entryDepositA := entries.Entry{AccountID: accountID_A, Amount: 15.00, EntryType: entries.EntryTypeDeposity}
	_, err = accountTransactionService.DepositAmount(entryDepositA)
	t.NoError(err)
	entryDepositB := entries.Entry{AccountID: accountID_B, Amount: 0.00, EntryType: entries.EntryTypeDeposity}
	_, err = accountTransactionService.DepositAmount(entryDepositB)
	t.NoError(err)

	// Act
	transactions, err := accountTransactionService.TransferAmount(accountID_A, accountID_B, transferAmount)

	// Assert
	t.NoError(err)
	t.Equal(accountID_A, transactions[0].Balance.ID)
	t.Equal(expectedBalanceA, transactions[0].Balance.Balance)
	t.Equal(expectedTypeA, transactions[0].Type)

	t.Equal(accountID_B, transactions[1].Balance.ID)
	t.Equal(expectedBalanceB, transactions[1].Balance.Balance)
	t.Equal(expectedTypeB, transactions[1].Type)
}
