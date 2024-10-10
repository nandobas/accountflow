package accountTransactionHandler

import (
	"accountflow/api/handlers"
	"accountflow/api/middlewares"
	accounttransactionmiddleware "accountflow/api/middlewares/accountTransactionMiddleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Event(c *gin.Context) {
	req := EventRequest{}
	var err error
	var response *middlewares.Response_t
	var fromAccountID, toAccountID int

	if err := c.BindJSON(&req); err != nil {
		response := middlewares.RetFail("cannot unmarshal JSON")
		handlers.RetFail(c, response)
		return
	}

	if req.FromAccountID != "" {

		fromAccountID, err = strconv.Atoi(req.FromAccountID)
		if err != nil || fromAccountID == 0 {
			response := middlewares.RetFail("invalid url param from account id must be integer")
			handlers.RetFail(c, response)
			return
		}
	}

	if req.ToAccountID != "" {

		toAccountID, err = strconv.Atoi(req.ToAccountID)
		if err != nil || toAccountID == 0 {
			response := middlewares.RetFail("invalid url param to account id must be integer")
			handlers.RetFail(c, response)
			return
		}
	}

	accounttransactionMiddleware := accounttransactionmiddleware.NewEvent(int64(fromAccountID), int64(toAccountID), req.Amount)

	switch req.EventType {
	case "deposit":
		response = accounttransactionMiddleware.Deposit()
	case "withdraw":
		response = accounttransactionMiddleware.Withdraw()
	case "transfer":
		response = accounttransactionMiddleware.Transfer()
	}

	if !response.Success {
		handlers.RetFail(c, response)
		return
	}

	handlers.ResponseCreated(c, response)
}
