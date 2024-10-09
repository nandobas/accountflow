package entries_test

import (
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

func (t *testServiceSuite) TestService_WhenGetBalanceFromNonExistingAccount_ExpectedError() {
	// Arrange
	accountID := int64(1234)
	expectedBalance := float64(0)
	entryRepository := repository.NewRepositoryCache()
	entryService := entries.NewService(entryRepository)

	// Act
	balance, err := entryService.GetBalanceByAccountID(accountID)

	// Assert
	t.Error(err)
	t.Equal(expectedBalance, balance)
}

func (t *testServiceSuite) TestService_WhenAppendEntry_ExpectedBalance() {
	// Arrange
	accountID := int64(100)
	expectedBalance := float64(11.90)
	entryRepository := repository.NewRepositoryCache()
	entryService := entries.NewService(entryRepository)

	entry1 := entries.Entry{
		ID:        1,
		AccountID: accountID,
		Amount:    10.00,
		EntryType: entries.EntryTypeDeposity,
	}

	entry2 := entries.Entry{
		ID:        2,
		AccountID: accountID,
		Amount:    3.10,
		EntryType: entries.EntryTypeWithdrawal,
	}

	entry3 := entries.Entry{
		ID:        3,
		AccountID: accountID,
		Amount:    5.00,
		EntryType: entries.EntryTypeDeposity,
	}

	// Act
	err := entryService.AppendEntry(entry1)
	t.NoError(err)
	err = entryService.AppendEntry(entry2)
	t.NoError(err)
	err = entryService.AppendEntry(entry3)
	t.NoError(err)

	// Assert
	t.NoError(err)
	balance, err := entryService.GetBalanceByAccountID(accountID)
	t.NoError(err)
	t.Equal(expectedBalance, balance)
}
