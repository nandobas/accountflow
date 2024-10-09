package api

import (
	"accountflow/api/handlers"
	"accountflow/api/handlers/accountTransaction"
)

func (s *Service) GetRoutes() {

	/* WEB ROUTES */

	s.Engine.GET("/test", handlers.TestHandle)
	s.Engine.POST("/reset", handlers.Reset)
	s.Engine.GET("/balance", accountTransaction.Balance)
	s.Engine.POST("/event", accountTransaction.Event)
}
