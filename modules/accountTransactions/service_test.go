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

	entry1 := entries.Entry{
		ID:        1,
		AccountID: accountID,
		Amount:    10.00,
		EntryType: entries.EntryTypeDeposity,
	}

	// Act
	transaction, err := accountTransactionService.DeposityAmount(entry1)

	// Assert
	t.NoError(err)
	t.Equal(expectedBalance, transaction.Balance)
	t.Equal(expectedType, transaction.Type)
}
