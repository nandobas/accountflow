package api

import (
	"accountflow/environment"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Service struct {
	*gin.Engine
}

func NewService() *Service {
	return &Service{
		gin.Default(),
	}
}

func (s *Service) Config() {

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowMethods("GET", "PUT", "POST", "DELETE", "OPTIONS")
	config.AddAllowHeaders("Content-Type", "Content-Length", "Authorization", "Origin", "Cache-Control", "Access-Control-Allow-Origin")
	s.Engine.Use(cors.New(config))
	s.Engine.Use(gin.Logger())
	s.Engine.Use(gin.Recovery())
	s.Engine.RedirectTrailingSlash = false
	s.Engine.RedirectFixedPath = true

	s.GetRoutes()

}

func (s *Service) Start() {
	s.Config()
	srv := &http.Server{
		Addr:    ":" + environment.ListenHttpPort,
		Handler: s.Engine,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// IdleTimeout:    30 * time.Second, //TODO: check here why is not given time between requisitions
		MaxHeaderBytes: 1 << 20,
	}
	srv.ListenAndServe()
}
