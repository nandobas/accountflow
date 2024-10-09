package repository_test

import (
	"accountflow/modules/entries"
	"accountflow/modules/entries/repository"
	"accountflow/modules/system/lcache"
	"testing"

	"github.com/stretchr/testify/suite"
)

type testRepositorySuite struct {
	suite.Suite
}

func TestService(t *testing.T) {
	suite.Run(t, new(testRepositorySuite))
}

func (t *testRepositorySuite) SetupTest() {

	lcache.InitLocalCache()
}

func (t *testRepositorySuite) TestRepository_WhenGetEntriesFromNonExistingAccount_ExpectedError() {
	//Arrange
	accountID := int64(1)

	r := repository.NewRepositoryCache()

	//Action
	entries, err := r.GetEntriesByAccountID(accountID)

	//Assert
	t.Error(err)
	t.Len(entries, 0)

}

func (t *testRepositorySuite) TestRepository_WhenGetEntriesFromExistingAccount_ExpectedEntries() {
	//Arrange
	accountID := int64(1)

	r := repository.NewRepositoryCache()

	r.AppendEntry(entries.Entry{
		ID:        1,
		AccountID: accountID,
		Amount:    0.001,
		EntryType: entries.EntryTypeDeposity,
	})

	//Action
	entries, err := r.GetEntriesByAccountID(accountID)

	//Assert
	t.NoError(err)
	t.Len(entries, 1)

}
