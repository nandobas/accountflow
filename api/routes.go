package api

import (
	"accountflow/api/handlers"
	"accountflow/api/handlers/accountTransactionHandler"
)

func (s *Service) GetRoutes() {

	/* WEB ROUTES */

	s.Engine.GET("/test", handlers.TestHandle)
	s.Engine.POST("/reset", handlers.Reset)
	s.Engine.GET("/balance", accountTransactionHandler.GetBalance)
	s.Engine.POST("/event", accountTransactionHandler.Event)
}
