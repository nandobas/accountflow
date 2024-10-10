package handlers

import (
	"accountflow/api/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, resp *middlewares.Response_t) {
	buildResponse(c, resp)
}

func ResponseCreated(c *gin.Context, resp *middlewares.Response_t) {
	buildCreatedResponse(c, resp)
}

func ResponseOk(c *gin.Context, resp string) {
	c.String(http.StatusOK, resp)
}

func RetFail(c *gin.Context, ret *middlewares.Response_t) {
	c.String(http.StatusNotFound, ret.Data.(string))
}

func buildResponse(c *gin.Context, resp *middlewares.Response_t) {

	status := http.StatusOK
	if !resp.Success {
		status = http.StatusBadRequest
	}
	c.JSON(status, resp)
}

func buildCreatedResponse(c *gin.Context, resp *middlewares.Response_t) {

	//{"success":true,"data":{"Type":"origin","Balance":{"ID":100,"Balance":10}}}

	//"destination": {"id": "100", "balance": 10}

	//tResponse := transactionResponse_t{Type: resp.Data}

	c.JSON(http.StatusCreated, resp.Data)
}
