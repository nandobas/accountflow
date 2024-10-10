package accountTransactionHandler

import (
	"accountflow/api/handlers"
	"accountflow/api/middlewares"
	accounttransactionmiddleware "accountflow/api/middlewares/accountTransactionMiddleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {

	acID := c.DefaultQuery("account_id", "")

	accountID, err := strconv.Atoi(acID)
	if err != nil || acID == "" || accountID == 0 {
		response := middlewares.RetFail("invalid url param account id must be integer")
		handlers.RetFail(c, response)
		return
	}

	accounttransactionMiddleware := accounttransactionmiddleware.NewBalance(int64(accountID))
	response := accounttransactionMiddleware.GetBalance()

	if !response.Success {

		handlers.RetFail(c, response)
		return
	}

	handlers.Response(c, response)
}
