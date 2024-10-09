package handlers

import (
	"accountflow/api/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, resp *middlewares.Response_t) {
	buildResponse(c, resp)
}

func RetFail(c *gin.Context, ret string) {
	buildResponse(c, middlewares.RetFail(ret))
}

func buildResponse(c *gin.Context, resp *middlewares.Response_t) {

	status := http.StatusOK
	if !resp.Success {
		status = http.StatusBadRequest
	}
	c.JSON(status, resp)
}
