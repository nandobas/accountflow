package accountTransaction

import (
	"accountflow/api/handlers"
	"accountflow/api/middlewares"

	"github.com/gin-gonic/gin"
)

func Balance(c *gin.Context) {

	response := middlewares.NewResponseObject(true, "Hello!")

	handlers.Response(c, response)
}
