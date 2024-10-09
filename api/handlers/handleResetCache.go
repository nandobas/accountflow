package handlers

import (
	"accountflow/api/middlewares"
	"accountflow/modules/system/lcache"

	"github.com/gin-gonic/gin"
)

func Reset(c *gin.Context) {

	var itens []int64
	allEntries := lcache.GetAllEntries()
	if allEntries != nil {
		itens = allEntries.([]int64)
	}
	lcache.DeleteAccountInfoItens(itens)
	response := middlewares.RetOK()

	ResponseOk(c, response)
}
