package handlers

import (
	"accountflow/api/middlewares"

	"github.com/gin-gonic/gin"
)

func TestHandle(c *gin.Context) {

	response := middlewares.NewResponseObject(true, "Hello!")

	Response(c, response)
}
