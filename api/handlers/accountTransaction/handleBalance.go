package accountTransaction

import (
	"accountflow/api/handlers"
	"accountflow/api/middlewares"

	"github.com/gin-gonic/gin"
)

func Balance(c *gin.Context) {

	response := middlewares.RetFail("0")

	handlers.RetFail(c, response)
}
