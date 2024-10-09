package handlers

import (
	"accountflow/api/middlewares"
	"accountflow/modules/system/lcache"

	"github.com/gin-gonic/gin"
)

func Reset(c *gin.Context) {

	lcache.Cleanner()
	response := middlewares.RetOK()

	Response(c, response)
}
