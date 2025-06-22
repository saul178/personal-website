package server

import (
	"github.com/gin-gonic/gin"
	"github.com/saul178/personal-website/middleware"
)

func NewServer() *gin.Engine {
	middleware.InfoLog.Println("initializing gin server")
	router := gin.Default()

	// can have middleware here to be used
	// like a logger, recovery, etc.
	middleware.InfoLog.Println("initializing cors middleware")
	router.Use(middleware.CORSMiddleware())

	return router
}
