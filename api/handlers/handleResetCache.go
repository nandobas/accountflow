package handlers

import (
	"accountflow/api/middlewares"
	"accountflow/modules/entries/repository"
	"accountflow/modules/system/lcache"

	"github.com/gin-gonic/gin"
)

func Reset(c *gin.Context) {

	var itens []int64
	allEntries := lcache.GetAllEntries()
	if allEntries != nil {
		givenUnitsFromCache := allEntries.([]lcache.UnitCache_t)

		for _, givenUnitCache := range givenUnitsFromCache {
			givenEntry := givenUnitCache.Item.(repository.Entry)
			itens = append(itens, givenEntry.ID)
		}
	}

	lcache.DeleteAccountInfoItens(itens)
	response := middlewares.RetOK()

	ResponseOk(c, response)
}
