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
	c.JSON(status, resp.Data)
}

func buildCreatedResponse(c *gin.Context, resp *middlewares.Response_t) {

	c.JSON(http.StatusCreated, resp.Data)
}
