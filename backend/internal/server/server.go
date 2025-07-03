package server

import (
	"github.com/gin-gonic/gin"
	"github.com/saul178/personal-website/config"
	"github.com/saul178/personal-website/middleware"
)

func NewServer(cfg *config.Config) *gin.Engine {
	middleware.Logger.Info("initializing gin server", "port", 8080)
	router := gin.Default()

	// can have middleware here to be used
	// like a logger, recovery, etc.
	router.Use(middleware.CORSMiddleware(cfg.Domain.DomainName))
	router.Use(middleware.CacheMiddleware())

	return router
}
