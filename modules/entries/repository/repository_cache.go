package repository

import (
	"accountflow/modules/system/lcache"
	"fmt"
)

func NewRepositoryCache() Repository {
	return &repositoryCache{}
}

type repositoryCache struct {
	entryIdx int64
}

func (r *repositoryCache) GetEntriesByAccountID(accountID int64) ([]Entry, error) {

	var response []Entry

	givenFromCache := lcache.GetAllEntries()
	if givenFromCache == nil {
		return response, fmt.Errorf("unable to get entries: empty repository")
	}

	givenUnitsFromCache := givenFromCache.([]lcache.UnitCache_t)
	for _, givenUnitCache := range givenUnitsFromCache {

		givenEntry := givenUnitCache.Item.(Entry)
		if givenEntry.AccountID == accountID {
			response = append(response, givenEntry)
		}
	}

	return response, nil
}

func (r *repositoryCache) AppendEntry(entry Entry) error {
	r.entryIdx++
	lcache.SetAccountEntryInfoCache(r.entryIdx, entry)
	return nil
}
