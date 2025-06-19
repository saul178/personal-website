package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/saul178/personal-website/middleware"
)

func NewServer() *gin.Engine {
	router := gin.Default()

	// can have middleware here to be used
	// like a logger, recovery, etc.
	fmt.Println("in middleware section before")
	router.Use(middleware.CORSMiddleware())
	fmt.Println("in middleware section after")

	return router
}
