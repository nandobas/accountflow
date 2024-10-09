package repository

import (
	"accountflow/modules/entries"
	"accountflow/modules/system/lcache"
	"fmt"
)

func NewRepositoryCache() Repository {
	return &repositoryCache{}
}

type repositoryCache struct {
}

func (r *repositoryCache) GetEntriesByAccountID(accountID int64) ([]entries.Entry, error) {

	var response []entries.Entry

	givenFromCache := lcache.GetAllEntries()
	if givenFromCache == nil {
		return response, fmt.Errorf("unable to get entries: empty repository")
	}

	givenUnitsFromCache := givenFromCache.([]lcache.UnitCache_t)
	for _, givenUnitCache := range givenUnitsFromCache {

		givenEntry := givenUnitCache.Item.(entries.Entry)
		if givenEntry.AccountID == accountID {
			response = append(response, givenEntry)
		}
	}

	if len(response) == 0 {
		return response, fmt.Errorf("unable to get entries from account id: %d", accountID)
	}

	return response, nil
}

func (r *repositoryCache) AppendEntry(entry entries.Entry) error {
	lcache.SetAccountEntryInfoCache(entry.ID, entry)
	return nil
}
