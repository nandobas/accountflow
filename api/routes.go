package api

import "accountflow/api/handlers"

func (s *Service) GetRoutes() {

	/* WEB ROUTES */

	s.Engine.GET("/test", handlers.TestHandle)
}
