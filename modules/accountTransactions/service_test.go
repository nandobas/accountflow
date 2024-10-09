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
	builderEntry := entries.StorageEntries{}
	accountTransactionService := accounttransactions.NewAccountTransactionService(entryService)

	entry := builderEntry.NewEntry(accountID, 15.00, entries.EntryTypeDeposity)

	// Act
	transaction, err := accountTransactionService.DeposityAmount(entry)

	// Assert
	t.NoError(err)
	t.Equal(accountID, transaction.Balance.ID)
	t.Equal(expectedBalance, transaction.Balance.Balance)
	t.Equal(expectedType, transaction.Type)
}
